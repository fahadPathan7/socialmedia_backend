package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct def
type React struct {
	ID        primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	PostID    primitive.ObjectID `json:"post_id" bson:"post_id"`
	Author	  string             `json:"author" bson:"author"`
	Type      string             `json:"type" bson:"type"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}