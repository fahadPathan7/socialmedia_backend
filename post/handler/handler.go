package handler

import (
	"context"

	// "github.com/fahadPathan7/socialmedia_backend/batman"
	"github.com/fahadPathan7/socialmedia_backend/batman/auth"
	"github.com/fahadPathan7/socialmedia_backend/batman/client"
	accesscontrol "github.com/fahadPathan7/socialmedia_backend/batman/access_control"
	"github.com/fahadPathan7/socialmedia_backend/post/model"
	"github.com/fahadPathan7/socialmedia_backend/post/repository"
	"github.com/fahadPathan7/socialmedia_backend/post/validation"
	pb "github.com/fahadPathan7/socialmedia_backend/proto/post"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// service struct def
type service struct {
	repo repository.PostRepository
	auth auth.Authenticator
	validator validation.RequestValidator
	pb.UnimplementedPostServiceServer // Embed the UnimplementedPostServiceServer
}

// create a new post
func (s *service) Create(ctx context.Context, r *pb.CreateRequest) (*pb.CreateResponse, error) {
	// // Validate the request
	// err := s.validator.ValidateCreateRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// check if the user from context is the user who created the post
	user, err := s.auth.UserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "User not found")
	}
	if user.Email != r.GetAuthor() {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to create post")
	}

	// Create a new post
	post := &model.Post{
		ID:      primitive.NewObjectID(), // it will generate a new object id for the post
		Title:   r.GetTitle(),
		Content: r.GetContent(),
		Author:  r.GetAuthor(),
	} // CreatedAt and UpdatedAt will be set to the current time by default

	// Save the post
	err = s.repo.Create(post)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create post")
	}

	// Return the response
	return &pb.CreateResponse{
		Id: post.ID.Hex(),
		Title: post.Title,
		Content: post.Content,
		Author: post.Author,
		CreatedAt: post.CreatedAt.String(),
		UpdatedAt: post.UpdatedAt.String(),
		Status: &pb.Status{Code: 201, Message: "Post created successfully"},
	}, nil
}

// get a post by id
func (s *service) Read(ctx context.Context, r *pb.ReadRequest) (*pb.ReadResponse, error) {
	// Validate the request
	// err := s.validator.ValidateGetRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get the post
	post, err := s.repo.GetByID(r.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Post not found")
	}

	// Return the response
	return &pb.ReadResponse{
		Id: post.ID.Hex(),
		Title: post.Title,
		Content: post.Content,
		Author: post.Author,
		CreatedAt: post.CreatedAt.String(),
		UpdatedAt: post.UpdatedAt.String(),
		Status: &pb.Status{Code: 200, Message: "Post found"},
	}, nil
}

// get all posts
func (s *service) ReadAll(ctx context.Context, r *pb.Empty) (*pb.ReadAllResponse, error) {
	// Validate the request
	// err := s.validator.ValidateGetAllRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// Get all posts
	posts, err := s.repo.GetAll()
	if err != nil {
		return nil, status.Error(codes.NotFound, "Posts not found")
	}

	// Create a new response
	response := &pb.ReadAllResponse{
		Posts: make([]*pb.Post, 0),
		Status: &pb.Status{Code: 200, Message: "Posts found"},
	}

	// Convert the posts to pb.Post
	for _, post := range posts {
		response.Posts = append(response.Posts, &pb.Post{
			Id: post.ID.Hex(),
			Title: post.Title,
			Content: post.Content,
			Author: post.Author,
			CreatedAt: post.CreatedAt.String(),
			UpdatedAt: post.UpdatedAt.String(),
		})
	}

	// Return the response
	return response, nil
}

// update a post
func (s *service) Update(ctx context.Context, r *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// // Validate the request
	// err := s.validator.ValidateUpdateRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// check if the user from context is the user who created the post
	user, err := s.auth.UserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "User not found")
	}

	// check if the user is the author of the post
	post, err := s.repo.GetByID(r.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Post not found")
	}
	if post.Author != user.Email {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to update post")
	}

	// convert the id string to an object id
	id, err := primitive.ObjectIDFromHex(r.GetId())

	// Create a new post
	post = &model.Post{
		ID:      id,
		Title:   r.GetTitle(),
		Content: r.GetContent(),
	}

	// Update the post
	err = s.repo.Update(post)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update post")
	}

	// Return the response
	return &pb.UpdateResponse{
		Id: r.GetId(),
		Title: post.Title,
		Content: post.Content,
		Author: post.Author,
		CreatedAt: post.CreatedAt.String(),
		UpdatedAt: post.UpdatedAt.String(),
		Status: &pb.Status{Code: 200, Message: "Post updated successfully"},
	}, nil
}

// delete a post
func (s *service) Delete(ctx context.Context, r *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	// Validate the request
	// err := s.validator.ValidateDeleteRequest(r)
	// if err != nil {
	// 	return nil, status.Error(codes.InvalidArgument, "Validation failed")
	// }

	// check if the user from context is the user who created the post or an admin
	user, err := s.auth.UserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "User not found")
	}

	// check if user roles include admin
	adminRole := false
	for _, role := range user.Roles {
		if role == accesscontrol.Admin {
			adminRole = true
			break
		}
	}

	// get the post from the database
	post, err := s.repo.GetByID(r.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "Post not found")
	}

	// check if the user is the author of the post or an admin
	if user.Email != post.Author && !adminRole {
		return nil, status.Error(codes.PermissionDenied, "User not allowed to delete post")
	}

	// Delete the post
	err = s.repo.Delete(r.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to delete post")
	}

	// Delete all the comments of the post
	err = client.DeleteAllCommentsOfAPost(r.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to delete post comments")
	}

	// Delete all the reactions to the post
	err = client.DeleteAllReactsOfAPost(r.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to delete post reactions")
	}

	// Return the response
	return &pb.DeleteResponse{
		Id: r.GetId(),
		Status: &pb.Status{Code: 200, Message: "Post deleted successfully"},
	}, nil
}

// NewService returns a new service
func NewService(repo repository.PostRepository, auth auth.Authenticator, validator validation.RequestValidator) pb.PostServiceServer {
	return &service{
		repo: repo,
		auth: auth,
		validator: validator,
	}
}