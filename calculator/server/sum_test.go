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
		expected     int32
		firstNumber  int32
		secondNumber int32
	}{
		{
			expected:     2,
			firstNumber:  1,
			secondNumber: 1,
		},
		{
			expected:     -1,
			firstNumber:  -2,
			secondNumber: 1,
		},
		{
			expected:     -1,
			firstNumber:  0,
			secondNumber: -1,
		},
	}

	for _, tt := range tests {
		req := &pb.SumRequest{FirstNumber: tt.firstNumber, SecondNumber: tt.secondNumber}
		res, err := c.Sum(context.Background(), req)

		if err != nil {
			t.Errorf("Sum(%v) got unexpected error", tt)
		}

		if res.Result != tt.expected {
			t.Errorf("SumResponse(%v, %v) = %v, expected %v", tt.firstNumber, tt.secondNumber, res.Result, tt.expected)
		}
	}
}
