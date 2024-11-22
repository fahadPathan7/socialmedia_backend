package repository

import (
	"github.com/fahadPathan7/socialmedia_backend/user/model"
)

type UserRepository interface {
	// Create a new user
	Create(user *model.User) error
	// Get a user by its ID
	GetByID(id string) (*model.User, error)
	// Get a user by its username
	GetByUsername(username string) (*model.User, error)
	// Get a user by its email
	GetByEmail(email string) (*model.User, error)
	// Update a user
	Update(user *model.User) error
	// Delete a user
	Delete(id string) error
	// Get all users
	GetAll() ([]*model.User, error)
}