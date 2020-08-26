package query

import (
	"context"
	"fmt"
	"log"

	"github.com/anshalshukla/folks/db/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// AddUser adds a new user to the database
func AddUser(user *models.User, client *mongo.Client) (string, error) {
	collection := client.Database("user_service").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	_id := fmt.Sprintf("%s", insertResult.InsertedID)

	_id = _id[10:34]
	return _id, nil
}
