package client

import (
	systemtoken "github.com/fahadPathan7/socialmedia_backend/batman/interceptors/system_token"
	"google.golang.org/grpc"
)

// GetConn returns a conn object
func getConn(systemCall bool, serverAddr string) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error
	if systemCall {
		conn, err = grpc.Dial(
			serverAddr,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(systemtoken.ClientInterceptor()),
		)
		if err != nil {
			return nil, err
		}
	} else {
		conn, err = grpc.Dial(
			serverAddr,
			grpc.WithInsecure(),
		)
		return nil, err

	}
	return conn, err
}
