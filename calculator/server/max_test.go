package main

import (
	"context"
	"io"
	"testing"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMax(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	stream, err := c.Max(context.Background())

	if err != nil {
		t.Fatalf("Error while opening stream and calling Max: %v", err)
	}

	waitc := make(chan struct{})
	errs := make(chan error, 1)
	var maximum int32

	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}

		for _, number := range numbers {
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				errs <- nil
				break
			}

			if err != nil {
				errs <- err
				break
			}

			maximum = res.Result
		}
		close(waitc)
	}()
	<-waitc

	err = <-errs

	if err != nil {
		t.Errorf("Error while receiving response: %v", err)
	}

	var expected int32 = 32

	if maximum != expected {
		t.Errorf("Expected max %v, got: %v", expected, maximum)
	}
}
