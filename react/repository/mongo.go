package repository

import (
	"context"
	"errors"
	"time"

	// pb "github.com/fahadPathan7/socialmedia_backend/proto/react"
	"github.com/fahadPathan7/socialmedia_backend/react/config"
	"github.com/fahadPathan7/socialmedia_backend/react/model"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// reactRepository struct
type reactRepository struct {
	client *mongo.Client
	config config.Config
}

// create
func (r *reactRepository) Create(react *model.React) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.ReactColl)
	_, err := coll.InsertOne(ctx, react)
	if err != nil {
		return errors.New("Failed to create react")
	}
	return nil
}

// get by id
func (r *reactRepository) GetByID(id string) (*model.React, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.ReactColl)
	// convert the id string to object id
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid id")
	}
	filter := bson.M{"_id": oid}
	var react model.React
	err = coll.FindOne(ctx, filter).Decode(&react)
	if err != nil {
		return nil, errors.New("React not found")
	}
	return &react, nil
}

// get by post id
func (r *reactRepository) GetByPostID(postID string) ([]*model.React, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.ReactColl)
	// convert the post id string to object id
	oid, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, errors.New("Invalid post id")
	}
	filter := bson.M{"post_id": oid}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, errors.New("Failed to get reacts")
	}
	var reacts []*model.React
	for cursor.Next(ctx) {
		var react model.React
		cursor.Decode(&react)
		reacts = append(reacts, &react)
	}
	return reacts, nil
}

// update
func (r *reactRepository) Update(react *model.React) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.ReactColl)
	// convert the id string to object id
	react.ID, _ = primitive.ObjectIDFromHex(react.ID.Hex())
	filter := bson.M{"_id": react.ID}
	update := bson.M{"$set": bson.M{"type": react.Type, "updated_at": time.Now()}}
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("Failed to update react")
	}
	return nil
}

// delete
func (r *reactRepository) Delete(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.ReactColl)
	// convert the id string to object id
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("Invalid id")
	}
	filter := bson.M{"_id": oid}
	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New("Failed to delete react")
	}
	return nil
}

// delete all reacts of a post
func (r *reactRepository) DeleteAll(postID string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.ReactColl)
	// convert the post id string to object id
	oid, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return errors.New("Invalid post id")
	}
	filter := bson.M{"post_id": oid}
	_, err = coll.DeleteMany(ctx, filter)
	if err != nil {
		return errors.New("Failed to delete reacts")
	}
	return nil
}

// NewMongoRepository creates a new react repository
func NewMongoRepository(dbclient *mongo.Client) ReactRepository {
	return &reactRepository{
		client: dbclient,
		config: config.New(),
	}
}