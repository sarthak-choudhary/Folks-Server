package query

import (
	"context"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllUsers - fetches all users from db
func GetAllUsers(client *mongo.Client) (interface{}, error) {
	var result models.Users
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("users")
	cur, err := collection.Find(ctx, bson.D{}, options.Find())

	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var user models.User
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return result, nil
}
