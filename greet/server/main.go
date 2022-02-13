package main

import (
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.GreetServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
