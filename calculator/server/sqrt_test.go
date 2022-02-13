package main

import (
	"context"
	"testing"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestSqrt(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: 25})

	var expected float64 = 5

	if res.Result != expected {
		t.Errorf("Expected %v, got: %v", expected, res.Result)
	}
}

func TestSqrtError(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	_, err = c.Sqrt(context.Background(), &pb.SqrtRequest{Number: -1})

	if err == nil {
		t.Error("Expected error, got nil")
	}

	respErr, ok := status.FromError(err)

	if !ok {
		t.Error("Expected error")
	}

	if respErr.Code() != codes.InvalidArgument {
		t.Errorf("Expected InvalidArgument, got %v", respErr.Code().String())
	}

	expected := "Received a negative number: -1"

	if respErr.Message() != expected {
		t.Errorf("Expected \"%s\", got %s", expected, respErr.Message())
	}
}
