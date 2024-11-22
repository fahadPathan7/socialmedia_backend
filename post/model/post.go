package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post struct def
type Post struct {
	ID        primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Author    string             `json:"author" bson:"author"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}