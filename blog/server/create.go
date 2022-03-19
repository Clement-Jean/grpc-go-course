package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog function was invoked with %v\n", in)

	data := BlogItem{
		AuthorID: in.GetAuthorId(),
		Title:    in.GetTitle(),
		Content:  in.GetContent(),
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	} else {
		return &pb.BlogId{
			Id: oid.Hex(),
		}, nil
	}
}
