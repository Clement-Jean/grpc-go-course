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

func TestRead(t *testing.T) {
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
		expectedBlog := BlogItem{
			ID:       primitive.NewObjectID(),
			AuthorID: "Clement",
			Title:    "A title",
			Content:  "Content !",
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedBlog.ID},
			{Key: "author_id", Value: expectedBlog.AuthorID},
			{Key: "title", Value: expectedBlog.Title},
			{Key: "content", Value: expectedBlog.Content},
		}))

		blogID := &pb.BlogId{
			Id: expectedBlog.ID.Hex(),
		}

		_, err := c.ReadBlog(context.Background(), blogID)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}

func TestReadError(t *testing.T) {
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
		mt.AddMockResponses(bson.D{{Key: "error", Value: 0}})

		blogID := &pb.BlogId{
			Id: primitive.NewObjectID().Hex(),
		}

		_, err := c.ReadBlog(context.Background(), blogID)

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

func TestReadInvalidIDError(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	blogID := &pb.BlogId{}

	_, err = c.ReadBlog(context.Background(), blogID)

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
