package repository

import (
	"context"
	"errors"
	"time"

	// pb "github.com/fahadPathan7/socialmedia_backend/proto/user"
	"github.com/fahadPathan7/socialmedia_backend/user/config"
	"github.com/fahadPathan7/socialmedia_backend/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// userRepository struct
type userRepository struct {
	client *mongo.Client
	config config.Config
}

// create
func (r *userRepository) Create(user *model.User) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.UserColl)
	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		return errors.New("Failed to create user")
	}
	return nil
}

// get by id
func (r *userRepository) GetByID(id string) (*model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.UserColl)
	// convert the id string to object id
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid id")
	}
	filter := bson.M{"_id": oid}
	var user model.User
	err = coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, errors.New("User not found")
	}
	return &user, nil
}

// get by username
func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.UserColl)
	filter := bson.M{"username": username}
	var user model.User
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, errors.New("User not found")
	}
	return &user, nil
}

// get by email
func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.UserColl)
	filter := bson.M{"email": email}
	var user model.User
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, errors.New("User not found")
	}
	return &user, nil
}

// update
func (r *userRepository) Update(user *model.User) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.UserColl)
	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": user}
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("Failed to update user")
	}
	return nil
}

// delete
func (r *userRepository) Delete(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.UserColl)
	filter := bson.M{"id": id}
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New("Failed to delete user")
	}
	return nil
}

// get all
func (r *userRepository) GetAll() ([]*model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	coll := r.client.Database(r.config.Database).Collection(r.config.UserColl)
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("Failed to get users")
	}
	defer cursor.Close(ctx)
	var users []*model.User
	for cursor.Next(ctx) {
		var user model.User
		cursor.Decode(&user)
		users = append(users, &user)
	}
	return users, nil
}

// NewMongoRepository returns a new user repo for mongo database
func NewMongoRepository(dbClient *mongo.Client) UserRepository {
	return &userRepository{
		client: dbClient,
		config: config.New(),
	}
}