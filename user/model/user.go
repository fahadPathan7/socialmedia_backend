package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct def
type User struct {
	ID        primitive.ObjectID    `json:"id" bson:"_id, omitempty"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	Email     string    `json:"email" bson:"email"`
    Roles	 []string   `json:"roles" bson:"roles"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}