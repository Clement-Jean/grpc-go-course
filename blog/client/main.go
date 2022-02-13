package main

import (
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}
