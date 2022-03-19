package main

import pb "github.com/Clement-Jean/grpc-go-course/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
