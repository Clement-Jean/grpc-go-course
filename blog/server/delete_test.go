package main

import (
	"context"
	"testing"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func TestDelete(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Success", func(mt *mtest.T) {
		collection = mt.Coll
		blogID := &pb.BlogId{
			Id: primitive.NewObjectID().Hex(),
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "acknowledged", Value: true},
			{Key: "n", Value: 1},
		})

		_, err := c.DeleteBlog(context.Background(), blogID)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}

func TestDeleteCountZeroError(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Error", func(mt *mtest.T) {
		collection = mt.Coll
		blogID := &pb.BlogId{
			Id: primitive.NewObjectID().Hex(),
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "acknowledged", Value: true},
			{Key: "n", Value: 0},
		})

		_, err := c.DeleteBlog(context.Background(), blogID)

		if err == nil {
			t.Error("Expected error")
		}

		e, ok := status.FromError(err)

		if !ok {
			t.Error("Expected error")
		}

		if e.Code() != codes.NotFound {
			t.Errorf("Expected NotFound, got %v", e.Code().String())
		}
	})
}

func TestDeleteOneError(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Error", func(mt *mtest.T) {
		collection = mt.Coll
		blogID := &pb.BlogId{
			Id: primitive.NewObjectID().Hex(),
		}
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{}))

		_, err := c.DeleteBlog(context.Background(), blogID)

		if err == nil {
			t.Error("Expected error")
		}

		e, ok := status.FromError(err)

		if !ok {
			t.Error("Expected error")
		}

		if e.Code() != codes.Internal {
			t.Errorf("Expected Internal, got %v", e.Code().String())
		}
	})
}

func TestDeleteInvalidIDError(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	blogID := &pb.BlogId{}

	_, err = c.DeleteBlog(context.Background(), blogID)

	if err == nil {
		t.Error("Expected error")
	}

	e, ok := status.FromError(err)

	if !ok {
		t.Error("Expected error")
	}

	if e.Code() != codes.InvalidArgument {
		t.Errorf("Expected InvalidArgument, got %v", e.Code().String())
	}
}
