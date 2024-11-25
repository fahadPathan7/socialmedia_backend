package client

import (
	// "context"
	"os"

	"github.com/fahadPathan7/socialmedia_backend/proto/react"
	// pb "github.com/fahadPathan7/socialmedia_backend/proto/react"
	"google.golang.org/grpc"
)

type reactClient struct{
	clent react.ReactServiceClient
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

