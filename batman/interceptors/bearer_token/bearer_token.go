package bearertoken

import (
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Authenticator interface def
type Authenticator interface {
	AuthGuarded(fullMethod string, notAuthGuardedEndpoints map[string]bool) (bool, error)
	Authenticate(ctx context.Context, req interface{}) (context.Context, error)
}

type authenticator struct{}

func (auth *authenticator) AuthGuarded(fullMethod string, notAuthGuardedEndpoints map[string]bool) (bool, error) {
	if notAuthGuardedEndpoints == nil {
		return true, fmt.Errorf("no settings found for auth guarded endpoints")
	}

	// Check if auth path is not auth guarded
	yes, ok := notAuthGuardedEndpoints[fullMethod]
	fmt.Println("auth guarded ", yes, ok)
	if ok && yes { // No need to validate token for this endpoint
		return false, nil
	}

	return true, nil
}

func (auth *authenticator) Authenticate(ctx context.Context, req interface{}) (context.Context, error) {
	// Extract headers
	var md metadata.MD
	var ok bool
	md, ok = metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "no headers found")
	}

	// Extract authorization header
	_, ok = md["authorization"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "authorization header not found")
	}

	return ctx, nil
}

// NewAuthenticator returns a new bearer token authenticator
func NewAuthenticator() Authenticator {
	return &authenticator{}
}

// Interceptor def
func Interceptor(bTokenAuth Authenticator, notAuthGuardedEndpoints map[string]bool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println(info.FullMethod)
		// Check if endpoint is auth guarded
		yes, err := bTokenAuth.AuthGuarded(info.FullMethod, notAuthGuardedEndpoints)
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("endpoint auth settings not found: %v", err))
		}

		if !yes {
			fmt.Println("no auth token required")
			return handler(ctx, req)
		}

		ctx, err = bTokenAuth.Authenticate(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("not authenticated: %v", err))
		}

		return handler(ctx, req)
	}
}

// ClientInterceptor intercepts all outgoing requests
func ClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("--> unary interceptor: %s", method)

		return invoker(attachToken(ctx), method, req, reply, cc, opts...)
	}
}

func attachToken(ctx context.Context) context.Context {
	fmt.Println("Adding token")
	// Extract headers
	var md metadata.MD
	var ok bool
	md, ok = metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("metadata not found")
		return ctx
	}

	// Extract authorization header
	bearerToken, ok := md["authorization"]
	if !ok {
		fmt.Println("token not found")
		return ctx
	}
	fmt.Println(strings.Join(bearerToken, " "))
	return metadata.AppendToOutgoingContext(ctx, "authorization", strings.Join(bearerToken, " "))
}
