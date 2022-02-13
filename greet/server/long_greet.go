package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func (*server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	result := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		firstName := req.GetFirstName()
		result += "Hello " + firstName + "!\n"
	}
}
