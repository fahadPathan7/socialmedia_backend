package repository

import (
	"github.com/fahadPathan7/socialmedia_backend/comment/model"
)

// CommentRepository interface
type CommentRepository interface {
	// Create a new comment
	CreateANewComment(comment *model.Comment) error
	// Get a comment by its ID
	GetCommentByID(id string) (*model.Comment, error)
	// Get all comments for a post
	GetAllCommentsForAPost(postID string) ([]*model.Comment, error)
	// Update a comment
	UpdateAComment(comment *model.Comment) error
	// Delete a comment
	DeleteAComment(id string) error
	// Delete all comments for a post
	DeleteAllCommentsForAPost(postID string) error
}

