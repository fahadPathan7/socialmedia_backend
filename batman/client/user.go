package client

import (
	"context"
	"os"

	"github.com/fahadPathan7/socialmedia_backend/proto/user"
	pb "github.com/fahadPathan7/socialmedia_backend/proto/user"
	"google.golang.org/grpc"
)

type userClient struct{}

// NewUserClient returns a new client for booking microservice
func NewUserClient(systemCall bool) (user.UserServiceClient, *grpc.ClientConn, error) {
	serverAddr := os.Getenv("USER_SERVICE_URI")
	conn, err := getConn(systemCall, serverAddr)
	if err != nil {
		return nil, conn, err
	}
	// defer conn.Close()
	c, err := user.NewUserServiceClient(conn), nil
	if err != nil {
		return nil, conn, err
	}
	return c, conn, nil
}

// get user by id
func (u *userClient) GetUserByID(ctx context.Context, id string) (*pb.User, error) {
	client, conn, err := NewUserClient(false)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	res, err := client.GetUserById(ctx, &pb.GetUserByIdRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return res.User, nil
}

// get user by email
func (u *userClient) GetUserByEmail(ctx context.Context, email string) (*pb.User, error) {
	client, conn, err := NewUserClient(false)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	res, err := client.GetUserByEmail(ctx, &pb.GetUserByEmailRequest{Email: email})
	if err != nil {
		return nil, err
	}

	return res.User, nil
}