package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (*server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")
	maximum := int32(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		number := req.GetNumber()

		if number > maximum {
			maximum = number
			sendErr := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
		}
	}
}
