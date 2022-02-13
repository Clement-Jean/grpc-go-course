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

func TestUpdate(t *testing.T) {
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
		expectedBlog := &blogItem{
			ID:       primitive.NewObjectID(),
			AuthorID: "not Clement",
			Title:    "a new Title",
			Content:  "a new content",
		}
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", bson.D{
				{"_id", expectedBlog.ID},
				{"author_id", expectedBlog.AuthorID},
				{"title", expectedBlog.Title},
				{"content", expectedBlog.Content},
			}},
		})

		_, err := c.UpdateBlog(context.Background(), documentToBlog(expectedBlog))

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}

func TestUpdateError(t *testing.T) {
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

		blog := &pb.Blog{
			Id:       primitive.NewObjectID().Hex(),
			AuthorId: "not Clement",
			Title:    "a new Title",
			Content:  "a new content",
		}

		_, err := c.UpdateBlog(context.Background(), blog)

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
