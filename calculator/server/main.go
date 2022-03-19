//go:build !test
// +build !test

package main

import (
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
