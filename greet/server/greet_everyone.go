package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func (*server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		firstName := req.GetFirstName()
		result := "Hello " + firstName + "!"

		sendErr := stream.Send(&pb.GreetResponse{
			Result: result,
		})

		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}
