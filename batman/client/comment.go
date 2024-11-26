package client

import (
	"context"
	"os"

	"github.com/fahadPathan7/socialmedia_backend/proto/comment"
	// pb "github.com/fahadPathan7/socialmedia_backend/proto/comment"
	"google.golang.org/grpc"
)

type commentClient struct{
	client comment.CommentServiceClient
}

// NewUserClient returns a new client for booking microservice
func NewCommentClient(systemCall bool) (comment.CommentServiceClient, *grpc.ClientConn, error) {
	serverAddr := os.Getenv("COMMENT_SERVICE_URI")
	conn, err := getConn(systemCall, serverAddr)
	if err != nil {
		return nil, conn, err
	}
	// defer conn.Close()
	c := comment.NewCommentServiceClient(conn)
	return c, conn, nil
}

// delete all comments for a post
func DeleteAllCommentsOfAPost(postID string) error {
	c, conn, err := NewCommentClient(true)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = c.DeleteAllCommentsOfAPost(context.Background(), &comment.DeleteAllRequest{
		PostId: postID,
	})
	if err != nil {
		return err
	}
	return nil
}

