package client

import (
	"context"
	"os"

	"github.com/fahadPathan7/socialmedia_backend/proto/react"
	// pb "github.com/fahadPathan7/socialmedia_backend/proto/react"
	"google.golang.org/grpc"
)

type reactClient struct{
	client react.ReactServiceClient
}

// NewUserClient returns a new client for booking microservice
func NewReactClient(systemCall bool) (react.ReactServiceClient, *grpc.ClientConn, error) {
	serverAddr := os.Getenv("REACT_SERVICE_URI")
	conn, err := getConn(systemCall, serverAddr)
	if err != nil {
		return nil, conn, err
	}
	// defer conn.Close()
	c := react.NewReactServiceClient(conn)
	return c, conn, nil
}

// delete all reacts for a post
func DeleteAllReactsOfAPost(postID string) error {
	c, conn, err := NewReactClient(true)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = c.DeleteAllReactsOfAPost(context.Background(), &react.DeleteAllReactsOfAPostRequest{
		PostId: postID,
	})
	if err != nil {
		return err
	}
	return nil
}

