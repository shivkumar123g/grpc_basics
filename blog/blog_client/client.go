package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shivkumar123g/grpc_go_course/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	certFile := "ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalf("Failed loading certificates: %v", sslErr)
		return
	}

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := blogpb.NewBlogServiceClient(cc)
	readBlog(c)

}

func createBlog(c blogpb.BlogServiceClient) {
	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AutherId: "1",
			Title:    "Title 3",
			Content:  "Content 3",
		},
	})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Println(res)
}

func readBlog(c blogpb.BlogServiceClient) {
	res, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "608236bcf26ded5be685f4a8",//608236bcf26ded5be685f4a8
	})
	if err != nil {
		fmt.Printf("ERROR: can't read: %v",err)
	}
	fmt.Println(res)
}
