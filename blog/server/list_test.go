package main

import (
	"context"
	"io"
	"testing"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestList(t *testing.T) {
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
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{"_id", primitive.NewObjectID()},
			{"author_id", "Clement"},
			{"title", "a title"},
			{"content", "a content"},
		})
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{"_id", primitive.NewObjectID()},
			{"author_id", "not Clement"},
			{"title", "another title"},
			{"content", "another content"},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		count := 0

		for {
			_, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			count++
		}

		if count != 2 {
			t.Errorf("Expected 2, got: %d", count)
		}
	})
}
