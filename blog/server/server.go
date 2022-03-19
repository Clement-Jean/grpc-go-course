package main

import pb "github.com/Clement-Jean/grpc-go-course/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
