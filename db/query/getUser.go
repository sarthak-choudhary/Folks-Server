package query

import (
	"context"

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
