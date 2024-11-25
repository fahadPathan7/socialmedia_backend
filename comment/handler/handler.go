package handler

import (
	"context"

	// "github.com/fahadPathan7/socialmedia_backend/batman"
	"github.com/fahadPathan7/socialmedia_backend/batman/auth"
	accesscontrol "github.com/fahadPathan7/socialmedia_backend/batman/access_control"
	"github.com/fahadPathan7/socialmedia_backend/batman/client"
	"github.com/fahadPathan7/socialmedia_backend/comment/model"
	"github.com/fahadPathan7/socialmedia_backend/comment/repository"
	"github.com/fahadPathan7/socialmedia_backend/comment/validation"
	pb "github.com/fahadPathan7/socialmedia_backend/proto/comment"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// service struct def
type service struct {
	repo repository.CommentRepository
	auth auth.Authenticator
	validator validation.RequestValidator
	pb.UnimplementedCommentServiceServer
}

// create a new comment
func (s *service) CreateComment(ctx context.Context, r *pb.CreateRequest) (*pb.CreateResponse, error) {
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

	// check if the user from context is the user who created the comment
	user, err := s.auth.UserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "User not found")
	}
	if user.Email != r.GetAuthor() {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to create comment")
	}

	// convert the post id string to object id
	postID, err := primitive.ObjectIDFromHex(r.GetPostId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid post id")
	}

	// Create a new comment
	comment := &model.Comment{
		ID:      primitive.NewObjectID(), // it will generate a new object id for the comment
		PostID:  postID, // converted to object id
		Content: r.GetContent(),
		Author:  r.GetAuthor(),
	}

	// Save the comment
	err = s.repo.CreateANewComment(comment)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create comment")
	}

	// Return the response
	return &pb.CreateResponse{
		Id: comment.ID.Hex(),
		PostId: comment.PostID.Hex(),
		Content: comment.Content,
		Author: comment.Author,
		CreatedAt: comment.CreatedAt.String(),
		UpdatedAt: comment.UpdatedAt.String(),
		Status: &pb.Status{Code: 201, Message: "Comment created successfully"},
	}, nil
}

// get a comment by id
func (s *service) ReadAComment(ctx context.Context, r *pb.ReadRequest) (*pb.ReadResponse, error) {
	// Validate the request
	// err := s.validator.ValidateGetRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get the comment
	comment, err := s.repo.GetCommentByID(r.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Comment not found")
	}

	// Return the response
	return &pb.ReadResponse{
		Id: comment.ID.Hex(),
		PostId: comment.PostID.Hex(),
		Content: comment.Content,
		Author: comment.Author,
		CreatedAt: comment.CreatedAt.String(),
		UpdatedAt: comment.UpdatedAt.String(),
		Status: &pb.Status{Code: 200, Message: "Comment found"},
	}, nil
}

// get all comments for a post
func (s *service) ReadAllCommentsOfAPost(ctx context.Context, r *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	// Validate the request
	// err := s.validator.ValidateGetAllRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get all comments for a post
	comments, err := s.repo.GetAllCommentsForAPost(r.GetPostId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Comments not found")
	}

	// Return the response
	var pbComments []*pb.Comment
	for _, comment := range comments {
		pbComments = append(pbComments, &pb.Comment{
			Id: comment.ID.Hex(),
			PostId: comment.PostID.Hex(),
			Content: comment.Content,
			Author: comment.Author,
			CreatedAt: comment.CreatedAt.String(),
			UpdatedAt: comment.UpdatedAt.String(),
		})
	}

	return &pb.ReadAllResponse{
		Comments: pbComments,
		Status: &pb.Status{Code: 200, Message: "Comments found"},
	}, nil
}

// update a comment
func (s *service) UpdateAComment(ctx context.Context, r *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// Validate the request
	// err := s.validator.ValidateUpdateRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// get the user from context (who is trying to update the comment)
	user, err := s.auth.UserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "User not found")
	}

	// Get the comment
	comment, err := s.repo.GetCommentByID(r.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Comment not found")
	}

	// check if the user from context is the user who created the comment
	if user.Email != comment.Author {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to update comment")
	}

	// Update the comment
	comment.Content = r.GetContent()
	err = s.repo.UpdateAComment(comment)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update comment")
	}

	// Return the response
	return &pb.UpdateResponse{
		Id: comment.ID.Hex(),
		PostId: comment.PostID.Hex(),
		Content: comment.Content,
		Author: comment.Author,
		CreatedAt: comment.CreatedAt.String(),
		UpdatedAt: comment.UpdatedAt.String(),
		Status: &pb.Status{Code: 200, Message: "Comment updated successfully"},
	}, nil
}

// delete a comment (user or admin)
func (s *service) DeleteAComment(ctx context.Context, r *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	// Validate the request
	// err := s.validator.ValidateDeleteRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// get the user from context (who is trying to delete the comment)
	user, err := s.auth.UserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "User not found")
	}

	// Get the comment
	comment, err := s.repo.GetCommentByID(r.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Comment not found")
	}

	// check if user roles include admin
	adminRole := false
	for _, role := range user.Roles {
		if role == accesscontrol.Admin {
			adminRole = true
			break
		}
	}

	// check if the user is the author of the comment or an admin
	if user.Email != comment.Author && !adminRole {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to delete comment")
	}

	// Delete the comment
	err = s.repo.DeleteAComment(r.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to delete comment")
	}

	// Return the response
	return &pb.DeleteResponse{
		Id: comment.ID.Hex(),
		Status: &pb.Status{Code: 200, Message: "Comment deleted successfully"},
	}, nil
}

// NewService creates a new comment service
func NewService(repo repository.CommentRepository, auth auth.Authenticator, validator validation.RequestValidator) pb.CommentServiceServer {
	return &service{
		repo: repo,
		auth: auth,
		validator: validator,
	}
}