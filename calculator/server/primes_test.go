package main

import (
	"context"
	"io"
	"reflect"
	"testing"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

func TestPrimes(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	req := &pb.PrimeRequest{
		Number: 567890,
	}

	res, err := c.Primes(context.Background(), req)

	if err != nil {
		t.Errorf("Primes(%v) got unexpected error", err)
	}

	count := 0
	primes := []int64{}

	for {
		prime, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Errorf("Error while reading stream: %v", err)
		}

		count++
		primes = append(primes, prime.Result)
	}

	if count != 4 {
		t.Errorf("Expected 4 elements, got: %d", count)
	}

	expected := []int64{2, 5, 109, 521}

	if !reflect.DeepEqual(primes, expected) {
		t.Errorf("Expected %v elements, got: %v", expected, primes)
	}
}
