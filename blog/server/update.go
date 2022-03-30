package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*pb.Blog, error) {
	log.Printf("UpdateBlog function was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	res := collection.FindOneAndUpdate(
		ctx,
		bson.D{{"_id", oid}},
		bson.D{{"$set", data}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	return documentToBlog(data), nil
}
