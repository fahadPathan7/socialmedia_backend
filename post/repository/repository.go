package repository

import (
	"github.com/fahadPathan7/socialmedia_backend/post/model"
)

// PostRepository interface
type PostRepository interface {
	// Create a new post
	Create(post *model.Post) error
	// Get a post by its ID
	GetByID(id string) (*model.Post, error)
	// Get all posts
	GetAll() ([]*model.Post, error)
	// Update a post
	Update(post *model.Post) error
	// Delete a post
	Delete(id string) error
}