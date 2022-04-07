package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func LogInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Printf("%s was invoked with %v\n", method, req)

		headers, ok := metadata.FromOutgoingContext(ctx)

		if ok {
			log.Printf("Sending headers: %v\n", headers)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
