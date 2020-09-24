package query

import (
	"context"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

//FollowUser - Logged in user follows the user with given username
func FollowUser(id primitive.ObjectID, username string, client *mongo.Client) (interface{}, error) {
	var err error
	var results models.User

	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id}
	q2 := bson.M{"$inc": bson.M{"followedByNumber": 1}}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("folks").Collection("events")
	err = collection.FindOneAndUpdate(ctx, q, q2).Decode(&results)
	if err != nil {
		return nil, err
	}

	q = bson.M{"username": username}
	q2 = bson.M{"$addToSet": bson.M{"followedBy": id}}
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection = client.Database("folks").Collection("users")
	err = collection.FindOneAndUpdate(ctx, q, q2).Err()
	if err != nil {
		return nil, err
	}
	return results, nil
}
