package repository

import (
	"context"
	"errors"
	"time"

	"github.com/fahadPathan7/socialmedia_backend/post/config"
	"github.com/fahadPathan7/socialmedia_backend/post/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// postRepository struct
type postRepository struct {
	client *mongo.Client
	config config.Config
}

// create a new post
func (r *postRepository) Create(post *model.Post) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.PostColl)
	_, err := coll.InsertOne(ctx, post)
	if err != nil {
		return errors.New("Failed to create post")
	}
	return nil
}

// get post by id
func (r *postRepository) GetByID(id string) (*model.Post, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.PostColl)
	// convert the id string to object id
	objectID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	var post model.Post
	err = coll.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return nil, errors.New("Post not found")
	}
	return &post, nil
}

// get all posts
func (r *postRepository) GetAll() ([]*model.Post, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.PostColl)
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("Failed to get posts")
	}
	var posts []*model.Post
	for cursor.Next(ctx) {
		var post model.Post
		cursor.Decode(&post)
		posts = append(posts, &post)
	}
	return posts, nil
}

// update a post
func (r *postRepository) Update(post *model.Post) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.PostColl)
	// convert the id string to object id
	post.ID, _ = primitive.ObjectIDFromHex(post.ID.Hex())
	filter := bson.M{"_id": post.ID}
	// update only those fields that are not empty
	update := bson.M{}
	if post.Title != "" {
		update["title"] = post.Title
	}
	if post.Content != "" {
		update["content"] = post.Content
	}
	update["updated_at"] = time.Now()
	_, err := coll.UpdateOne(ctx, filter , bson.M{"$set": update})
	if err != nil {
		return errors.New("Failed to update post")
	}
	return nil
}

// delete a post
func (r *postRepository) Delete(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.PostColl)
	// convert the id string to object id
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New("Failed to delete post")
	}
	return nil
}

// NewPostRepository returns a new post repository
func NewMongoRepository(dbclient *mongo.Client) PostRepository {
	return &postRepository{
		client: dbclient,
		config: config.New(),
	}
}