package main

import (
	"context"
	"io"
	"testing"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGreetManyTimes(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	req := &pb.GreetRequest{
		FirstName: "Clement",
	}

	res, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		t.Errorf("GreetManyTimes(%v) got unexpected error", err)
	}

	count := 0

	for {
		_, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Errorf("Error while reading stream: %v", err)
		}

		count++
	}

	if count != 10 {
		t.Errorf("Expected 10 elements, got: %d", count)
	}
}
