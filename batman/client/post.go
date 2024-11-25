package client

import (
	"context"
	"os"

	"github.com/fahadPathan7/socialmedia_backend/proto/post"
	pb "github.com/fahadPathan7/socialmedia_backend/proto/post"
	"google.golang.org/grpc"
)

type postClient struct{
	clent post.PostServiceClient
}

// NewUserClient returns a new client for booking microservice
func NewPostClient(systemCall bool) (post.PostServiceClient, *grpc.ClientConn, error) {
	serverAddr := os.Getenv("POST_SERVICE_URI")
	conn, err := getConn(systemCall, serverAddr)
	if err != nil {
		return nil, conn, err
	}
	// defer conn.Close()
	c := post.NewPostServiceClient(conn)
	return c, conn, nil
}

// get post by id
func GetPostByID(ctx context.Context, id string) (*pb.ReadResponse, error) {
	client, conn, err := NewPostClient(true)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	res, err := client.Read(ctx, &pb.ReadRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return res, nil
}

