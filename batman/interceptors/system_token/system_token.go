package systemtoken

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ClientInterceptor intercepts all outgoing requests and adds a system token which identifies that the request was made
// from a system and not a user
func ClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("Attaching system call token: %s", method)

		return invoker(attachToken(ctx), method, req, reply, cc, opts...)
	}
}

func attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "system-authorized", "true")
}
