package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (*Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")
	var maximum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		if number := req.Number; number > maximum {
			maximum = number
			sendErr := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
			}
		}
	}
}
