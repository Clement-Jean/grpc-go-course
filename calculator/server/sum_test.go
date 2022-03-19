package main

import (
	"context"
	"testing"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestSum(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	tests := []struct {
		expected      int32
		first_number  int32
		second_number int32
	}{
		{
			expected:      2,
			first_number:  1,
			second_number: 1,
		},
		{
			expected:      -1,
			first_number:  -2,
			second_number: 1,
		},
		{
			expected:      -1,
			first_number:  0,
			second_number: -1,
		},
	}

	for _, tt := range tests {
		req := &pb.SumRequest{FirstNumber: tt.first_number, SecondNumber: tt.second_number}
		res, err := c.Sum(context.Background(), req)

		if err != nil {
			t.Errorf("Sum(%v) got unexpected error", tt)
		}

		if res.Result != tt.expected {
			t.Errorf("SumResponse(%v, %v) = %v, expected %v", tt.first_number, tt.second_number, res.Result, tt.expected)
		}
	}
}
