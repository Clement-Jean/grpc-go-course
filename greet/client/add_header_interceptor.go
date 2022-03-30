package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AddHeaderInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		send, _ := metadata.FromOutgoingContext(ctx)
		newMD := metadata.Pairs("authorization", "aDummyToken")
		ctx = metadata.NewOutgoingContext(ctx, metadata.Join(send, newMD))

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
