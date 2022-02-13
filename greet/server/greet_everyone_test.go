package main

import (
	"context"
	"io"
	"testing"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func TestGreetEveryone(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		t.Errorf("Error while creating stream: %v", err)
	}

	requests := []*pb.GreetRequest{
		{
			FirstName: "Clement",
		},
		{
			FirstName: "Marie",
		},
		{
			FirstName: "Test",
		},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			stream.Send(req)
		}

		stream.CloseSend()
	}()

	go func() {
		idx := 0

		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				t.Errorf("Error while receiving: %v", err)
				break
			}

			expected := "Hello " + requests[idx].FirstName + "!"
			if resp.Result != expected {
				t.Errorf("Expected \"%s\", got: \"%s\"", expected, resp.Result)
			}

			idx++
		}

		close(waitc)
	}()

	<-waitc
}
