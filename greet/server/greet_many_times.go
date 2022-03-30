package main

import (
	"log"
	"strconv"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)

	firstName := in.FirstName

	for i := 0; i < 10; i++ {
		res := "Hello " + firstName + " number " + strconv.Itoa(i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
