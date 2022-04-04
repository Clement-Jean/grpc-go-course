package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("----readBlog has been invoked----")
	_, err2 := c.ReadBlog(context.Background(), &pb.BlogId{Id: "aNonExistingId"})

	if err2 != nil {
		log.Printf("Error happened while reading: %v\n", err2)
	}

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
