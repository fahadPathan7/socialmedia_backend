package repository

import (
	"github.com/fahadPathan7/socialmedia_backend/react/model"
)

// ReactRepository interface
type ReactRepository interface {
	// Create a new react
	Create(react *model.React) error
	// Get a react by its ID
	GetByID(id string) (*model.React, error)
	// Get all reacts of a post
	GetByPostID(postID string) ([]*model.React, error)
	// Update a react
	Update(react *model.React) error
	// Delete a react
	Delete(id string) error
	// Delete all reacts of a post
	DeleteAll(postID string) error
}