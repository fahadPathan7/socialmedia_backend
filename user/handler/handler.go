package handler

import (
	"context"

	"github.com/fahadPathan7/socialmedia_backend/batman/auth"
	pb "github.com/fahadPathan7/socialmedia_backend/proto/user"
	"github.com/fahadPathan7/socialmedia_backend/user/model"
	"github.com/fahadPathan7/socialmedia_backend/user/repository"
	"github.com/fahadPathan7/socialmedia_backend/user/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// service struct def
type service struct {
    repo      repository.UserRepository
    auth      auth.Authenticator
    validator validation.RequestValidator
    pb.UnimplementedUserServiceServer // Embed the UnimplementedUserServiceServer
}

// Register a new user
func (s *service) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// // Validate the request
	// err := s.validator.ValidateRegisterRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Check if the user already exists
	_, err := s.repo.GetByUsername(r.Username)
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "User already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to hash password")
	}

	// Create a new user
	user := &model.User{
		ID: 	  primitive.NewObjectID(), // it will generate a new object id for the user
		Username: r.Username,
		Password: string(hashedPassword),
		Email:    r.Email,
		Roles:    []string{"user"},
	}
	err = s.repo.Create(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create user")
	}

	// Encode the user
	token, err := s.auth.EncodeUser(pb.User{
		Id:       user.ID.Hex(), // convert the object id to a string
		Username: user.Username,
		Email:    user.Email,
		Roles:    user.Roles,
	}, 24)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to encode user")
	}

	return &pb.RegisterResponse{
		Id:    	  user.ID.Hex(), // convert the object id to a string
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
		Status:   &pb.Status{Code: 201, Message: "User created"},
	}, nil
}

// Login a user
func (s *service) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	// // Validate the request
	// err := s.validator.ValidateLoginRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get the user by username
	user, err := s.repo.GetByEmail(r.Email)
	if err != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	// revert the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid password")
	}

	// Encode the user
	token, err := s.auth.EncodeUser(pb.User{
		Id:       user.ID.Hex(), // convert the object id to a string
		Username: user.Username,
		Email:    user.Email,
		Roles:    user.Roles,
	}, 24)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to encode user")
	}

	return &pb.LoginResponse{
		Token:    token,
		Status:   &pb.Status{Code: 200, Message: "User logged in"},
	}, nil
}

// Get a user by id
func (s *service) GetUserByID(ctx context.Context, r *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	// // Validate the request
	// err := s.validator.ValidateGetUserByIDRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get the user by id
	user, err := s.repo.GetByID(r.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pb.GetUserByIdResponse{
		User: &pb.User{
			Id:       user.ID.Hex(), // convert the object id to a string
			Username: user.Username,
			Email:    user.Email,
			Roles:    user.Roles,
		},
		Status: &pb.Status{Code: 200, Message: "User found"},
	}, nil
}

// Get a user by email
func (s *service) GetUserByEmail(ctx context.Context, r *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
	// // Validate the request
	// err := s.validator.ValidateGetUserByEmailRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get the user by email
	user, err := s.repo.GetByEmail(r.GetEmail())
	if err != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pb.GetUserByEmailResponse{
		User: &pb.User{
			Id:       user.ID.Hex(), // convert the object id to a string
			Username: user.Username,
			Email:    user.Email,
			Roles:    user.Roles,
		},
		Status: &pb.Status{Code: 200, Message: "User found"},
	}, nil
}

// NewService returns a new user server
func NewService(
	repo repository.UserRepository,
	authenticator auth.Authenticator,
	validator validation.RequestValidator,
) pb.UserServiceServer {
	return &service{
		repo:      repo,
		auth:      authenticator,
		validator: validator,
	}
}