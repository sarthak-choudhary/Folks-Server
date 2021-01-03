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
	collection := client.Database("folks").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		fmt.Println("here")
		log.Fatal(err)
		return "", err
	}

	id := fmt.Sprintf("%s", insertResult.InsertedID)

	id = id[10:34]
	return id, nil
}
