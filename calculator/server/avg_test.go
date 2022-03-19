package main

import (
	"context"
	"testing"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestAvg(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	stream, err := c.Avg(context.Background())

	if err != nil {
		t.Fatalf("Error while opening stream: %v", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		t.Errorf("Error while receiving response: %v", err)
	}

	var expected float64 = (3 + 5 + 9 + 54 + 23) / float64(len(numbers))

	if res.Result != expected {
		t.Errorf("Expected avg of %v, got: %v", expected, res.Result)
	}
}
