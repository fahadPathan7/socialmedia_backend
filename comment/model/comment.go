package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comment struct def
type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	PostID    primitive.ObjectID `json:"post_id" bson:"post_id"`
	Content   string             `json:"content" bson:"content"`
	Author    string             `json:"author" bson:"author"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}