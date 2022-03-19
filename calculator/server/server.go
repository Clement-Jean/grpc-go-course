package main

import pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
