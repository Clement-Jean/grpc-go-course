package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGreetWithDeadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	req := &pb.GreetRequest{
		FirstName: "Clement",
		LastName:  "Jean",
	}
	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	expected := "Hello Clement"

	if res.Result != expected {
		t.Errorf("Expected %s, got: %s", expected, res.Result)
	}
}

func TestGreetWithDeadlineExceeded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	req := &pb.GreetRequest{
		FirstName: "Clement",
		LastName:  "Jean",
	}
	_, err = c.GreetWithDeadline(ctx, req)

	if err == nil {
		t.Error("Expected error, got: nil")
	}

	statusErr, ok := status.FromError(err)

	if !ok {
		t.Error("Expected ok")
	}

	if statusErr.Code() != codes.DeadlineExceeded {
		t.Errorf("Expected DeadlineExceeded, got: %s", statusErr.Code().String())
	}
}
