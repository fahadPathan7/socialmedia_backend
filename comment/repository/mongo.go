package repository

import (
	"context"
	"errors"
	"time"

	"github.com/fahadPathan7/socialmedia_backend/comment/config"
	"github.com/fahadPathan7/socialmedia_backend/comment/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// commentRepository struct
type commentRepository struct {
	client *mongo.Client // mongo client. the connection is from the main.go
	config config.Config
}

// CreateANewComment creates a new comment
func (r *commentRepository) CreateANewComment(comment *model.Comment) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.CommentColl)
	_, err := coll.InsertOne(ctx, comment)
	if err != nil {
		return errors.New("Failed to create comment")
	}
	return nil
}

// GetCommentByID gets a comment by its ID
func (r *commentRepository) GetCommentByID(id string) (*model.Comment, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.CommentColl)
	// convert the id string to object id
	objectID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	var comment model.Comment
	err = coll.FindOne(ctx, filter).Decode(&comment)
	if err != nil {
		return nil, errors.New("Comment not found")
	}
	return &comment, nil
}

// GetAllCommentsForAPost gets all comments for a post
func (r *commentRepository) GetAllCommentsForAPost(postID string) ([]*model.Comment, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.CommentColl)
	// convert the post id string to object id
	objectID, err := primitive.ObjectIDFromHex(postID)
	filter := bson.M{"post_id": objectID}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, errors.New("Failed to get comments")
	}
	var comments []*model.Comment
	for cursor.Next(ctx) {
		var comment model.Comment
		cursor.Decode(&comment)
		comments = append(comments, &comment)
	}
	return comments, nil
}

// UpdateAComment updates a comment
func (r *commentRepository) UpdateAComment(comment *model.Comment) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.CommentColl)
	// convert the id string to object id
	comment.ID, _ = primitive.ObjectIDFromHex(comment.ID.Hex())
	filter := bson.M{"_id": comment.ID}
	update := bson.M{"$set": bson.M{"content": comment.Content, "updated_at": time.Now()}}
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("Failed to update comment")
	}
	return nil
}

// DeleteAComment deletes a comment
func (r *commentRepository) DeleteAComment(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.CommentColl)
	// convert the id string to object id
	objectID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New("Failed to delete comment")
	}
	return nil
}

// DeleteAllCommentsForAPost deletes all comments for a post
func (r *commentRepository) DeleteAllCommentsForAPost(postID string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.CommentColl)
	// convert the post id string to object id
	objectID, err := primitive.ObjectIDFromHex(postID)
	filter := bson.M{"post_id": objectID}
	_, err = coll.DeleteMany(ctx, filter)
	if err != nil {
		return errors.New("Failed to delete comments")
	}
	return nil
}

// NewMongoRepository returns a new comment repository
func NewMongoRepository(dbclient *mongo.Client) CommentRepository {
	return &commentRepository{
		client: dbclient,
		config: config.New(),
	}
}
