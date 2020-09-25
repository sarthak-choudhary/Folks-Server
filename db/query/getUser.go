package query

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/anshalshukla/folks/db/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//GetUser queries the database and gets the user by email.
func GetUser(email string, client *mongo.Client) (*models.User, error) {
	var result models.User
	filter := bson.D{{"email", email}}
	collection := client.Database("folks").Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

//GetUserByID queries the database and gets the user by ID.
func GetUserByID(_id string, client *mongo.Client) (interface{}, error) {
	var result models.User
	var err error

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		return nil, err
	}

	q := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("users")
	err = collection.FindOne(ctx, q).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
