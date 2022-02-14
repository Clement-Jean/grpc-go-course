package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func LogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Printf("Received a request: %v\n", req)

		headers, ok := metadata.FromIncomingContext(ctx)

		if ok {
			log.Printf("Received headers: %v\n", headers)
		}

		return handler(ctx, req)
	}
}
