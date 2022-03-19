package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (*Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	sum := int32(0)
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			return stream.SendAndClose(&pb.AvgResponse{
				Result: average,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		sum += req.GetNumber()
		count++
	}
}
