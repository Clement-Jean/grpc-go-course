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
	"google.golang.org/grpc/status"
)

func TestRead(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Success", func(mt *mtest.T) {
		collection = mt.Coll
		expectedBlog := blogItem{
			ID:       primitive.NewObjectID(),
			AuthorID: "Clement",
			Title:    "A title",
			Content:  "Content !",
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{"_id", expectedBlog.ID},
			{"author_id", expectedBlog.AuthorID},
			{"title", expectedBlog.Title},
			{"content", expectedBlog.Content},
		}))

		blogId := &pb.BlogId{
			Id: expectedBlog.ID.Hex(),
		}

		_, err := c.ReadBlog(context.Background(), blogId)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}

func TestReadError(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Error", func(mt *mtest.T) {
		collection = mt.Coll
		mt.AddMockResponses(bson.D{{"error", 0}})

		blogId := &pb.BlogId{
			Id: primitive.NewObjectID().Hex(),
		}

		_, err := c.ReadBlog(context.Background(), blogId)

		if err == nil {
			t.Error("Expected error")
		}

		respErr, ok := status.FromError(err)

		if !ok {
			t.Error("Expected error")
		}

		if respErr.Code() != codes.NotFound {
			t.Errorf("Expected NotFound, got %v", respErr.Code().String())
		}
	})
}
