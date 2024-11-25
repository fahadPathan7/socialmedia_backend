package handler

import (
	"context"

	"github.com/fahadPathan7/socialmedia_backend/batman/auth"
	"github.com/fahadPathan7/socialmedia_backend/batman/client"
	pb "github.com/fahadPathan7/socialmedia_backend/proto/react"
	"github.com/fahadPathan7/socialmedia_backend/react/model"
	"github.com/fahadPathan7/socialmedia_backend/react/repository"
	"github.com/fahadPathan7/socialmedia_backend/react/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// service struct def
type service struct {
    repo      repository.ReactRepository
    auth      auth.Authenticator
    validator validation.RequestValidator
    pb.UnimplementedReactServiceServer
}

// create a new react
func (s *service) CreateAReact(ctx context.Context, r *pb.CreateRequest) (*pb.CreateResponse, error) {
	// // Validate the request
	// err := s.validator.ValidateCreateRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// check if the post exists
	_, err := client.GetPostByID(ctx, r.GetPostId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Post not found")
	}

	// check if the user from context is the user who created the react
	user, err := s.auth.UserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "User not found")
	}
	if user.Email != r.GetAuthor() {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to create react")
	}

	// convert the post id string to object id
	postIDO, err := primitive.ObjectIDFromHex(r.GetPostId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid post id")
	}

	// Create a new react
	react := &model.React{
		ID:      primitive.NewObjectID(), // it will generate a new object id for the react
		PostID:  postIDO, // converted to object id
		Author:  r.GetAuthor(),
		Type:    r.GetType(),
	}

	// Save the react
	err = s.repo.Create(react)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create react")
	}

	// Return the response
	return &pb.CreateResponse{
		Id: react.ID.Hex(),
		PostId: react.PostID.Hex(),
		Author: react.Author,
		Type: react.Type,
		CreatedAt: react.CreatedAt.String(),
		UpdatedAt: react.UpdatedAt.String(),
		Status: &pb.Status{Code: 201, Message: "React created successfully"},
	}, nil
}

// get a react by id
func (s *service) ReadAReact(ctx context.Context, r *pb.ReadRequest) (*pb.ReadResponse, error) {
	// Validate the request
	// err := s.validator.ValidateGetRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get the react
	react, err := s.repo.GetByID(r.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "React not found")
	}

	// Return the response
	return &pb.ReadResponse{
		Id: react.ID.Hex(),
		PostId: react.PostID.Hex(),
		Author: react.Author,
		Type: react.Type,
		CreatedAt: react.CreatedAt.String(),
		UpdatedAt: react.UpdatedAt.String(),
		Status: &pb.Status{Code: 200, Message: "React found"},
	}, nil
}

// get all reacts of a post
func (s *service) ReadAllReactsOfAPost(ctx context.Context, r *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	// Validate the request
	// err := s.validator.ValidateGetAllRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get the reacts
	reacts, err := s.repo.GetByPostID(r.GetPostId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Reacts not found")
	}

	// Return the response
	var reactsResp []*pb.React
	for _, react := range reacts {
		reactsResp = append(reactsResp, &pb.React{
			Id: react.ID.Hex(),
			PostId: react.PostID.Hex(),
			Author: react.Author,
			Type: react.Type,
			CreatedAt: react.CreatedAt.String(),
			UpdatedAt: react.UpdatedAt.String(),
		})
	}

	return &pb.ReadAllResponse{
		Reacts: reactsResp,
		Status: &pb.Status{Code: 200, Message: "Reacts found"},
	}, nil
}

// update a react
func (s *service) UpdateAReact(ctx context.Context, r *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// Validate the request
	// err := s.validator.ValidateUpdateRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get the react
	react, err := s.repo.GetByID(r.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "React not found")
	}

	// check if the user from context is the user who created the react
	user, err := s.auth.UserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "User not found")
	}
	if user.Email != react.Author {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to update react")
	}

	// Update the react
	react.Type = r.GetType()
	err = s.repo.Update(react)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update react")
	}

	// Return the response
	return &pb.UpdateResponse{
		Id: react.ID.Hex(),
		PostId: react.PostID.Hex(),
		Author: react.Author,
		Type: react.Type,
		CreatedAt: react.CreatedAt.String(),
		UpdatedAt: react.UpdatedAt.String(),
		Status: &pb.Status{Code: 200, Message: "React updated successfully"},
	}, nil
}

// delete a react
func (s *service) DeleteAReact(ctx context.Context, r *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	// Validate the request
	// err := s.validator.ValidateDeleteRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get the react
	react, err := s.repo.GetByID(r.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "React not found")
	}

	// check if the user from context is the user who created the react
	user, err := s.auth.UserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "User not found")
	}
	if user.Email != react.Author {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to delete react")
	}

	// Delete the react
	err = s.repo.Delete(r.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to delete react")
	}

	// Return the response
	return &pb.DeleteResponse{
		Status: &pb.Status{Code: 200, Message: "React deleted successfully"},
	}, nil
}

// DeleteAllReactsOfAPost deletes all reacts of a post
func (s *service) DeleteAllReactsOfAPost(ctx context.Context, r *pb.DeleteAllReactsOfAPostRequest) (*pb.DeleteAllReactsOfAPostResponse, error) {
	// Validate the request
	// err := s.validator.ValidateDeleteAllRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get the post
	_, err := client.GetPostByID(ctx, r.GetPostId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Post not found")
	}

	// check if the user from context is the user who created the post
	user, err := s.auth.UserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "User not found")
	}

	// check if the user is the author of the post
	post, err := client.GetPostByID(ctx, r.GetPostId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Post not found")
	}
	if post.GetAuthor() != user.Email {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to delete reacts")
	}

	// Delete all reacts of the post
	err = s.repo.DeleteAll(r.GetPostId())
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to delete reacts")
	}

	// Return the response
	return &pb.DeleteAllReactsOfAPostResponse{
		PostId: r.GetPostId(),
		Status: &pb.Status{Code: 200, Message: "Reacts deleted successfully"},
	}, nil
}

// NewService returns the instance of react service
func NewService(repo repository.ReactRepository, auth auth.Authenticator, validator validation.RequestValidator) pb.ReactServiceServer {
	return &service{
		repo: repo,
		auth: auth,
		validator: validator,
	}
}