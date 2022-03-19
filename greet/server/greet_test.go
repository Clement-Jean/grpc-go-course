package main

import (
	"context"
	"testing"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGreet(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	tests := []struct {
		expected string
		name     string
	}{
		{
			expected: "Hello Clement",
			name:     "Clement",
		},
		{
			expected: "Hello Marie",
			name:     "Marie",
		},
		{
			expected: "Hello Test",
			name:     "Test",
		},
	}

	for _, tt := range tests {
		req := &pb.GreetRequest{FirstName: tt.name}
		res, err := c.Greet(context.Background(), req)

		if err != nil {
			t.Errorf("Greet(%v) got unexpected error", tt)
		}

		if res.Result != tt.expected {
			t.Errorf("GreetResponse(%v) = %v, expected %v", tt.name, res.Result, tt.expected)
		}
	}
}
