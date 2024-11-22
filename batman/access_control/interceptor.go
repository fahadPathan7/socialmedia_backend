package accesscontrol

import (
	"context"
	"fmt"

	"github.com/fahadPathan7/socialmedia_backend/batman/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Interceptor def
func Interceptor(authenticator auth.Authenticator, notAuthGuardedEndpointsMap map[string]bool, accessableRolesMap map[string][]string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// Check if system generated call
		fmt.Println("inside access controll interceptor")
		yes := isSystemGeneratedCall(ctx)
		if yes { // System call, no validation needed
			return handler(ctx, req)
		}

		fmt.Println("is auth guarded")
		fmt.Println(info.FullMethod)
		// Check if auth path is not auth guarded
		yes, ok := notAuthGuardedEndpointsMap[info.FullMethod]
		if ok && yes { // No need to validate token for this endpoint
			return handler(ctx, req)
		}

		// Extract user from auth token
		user, err := authenticator.UserFromContext(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("not authenticated: %v", err))
		}

		// Check if the role has permission to this endpoint
		accessableRoles, ok := accessableRolesMap[info.FullMethod]
		if !ok {
			return nil, status.Errorf(codes.PermissionDenied, "Not found")
		}

		for _, role := range user.Roles {
			if stringInSlice(role, accessableRoles) {
				return handler(ctx, req)
			}
		}

		return nil, status.Errorf(codes.PermissionDenied, "Not found")
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isSystemGeneratedCall(ctx context.Context) bool {
	// Extract headers
	var md metadata.MD
	var ok bool
	md, ok = metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	fmt.Println("Extracting system token")
	// Extract system-authorized header
	_, ok = md["system-authorized"]
	if !ok {
		return false
	}

	return true
}
