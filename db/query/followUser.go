package query

import (
	"context"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

//FollowUser - Logged in user follows the user with given username
func FollowUser(id primitive.ObjectID, user_id string, client *mongo.Client) (interface{}, error) {
	var err error
	var results models.User

	userID, err := primitive.ObjectIDFromHex(user_id)

	if err != nil {
		return nil, err
	}

	q := bson.M{"_id": userID}
	q2 := bson.M{"$addToSet": bson.M{"requestsReceived": id}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	collection := client.Database("folks").Collection("users")
	err = collection.FindOneAndUpdate(ctx, q, q2).Err()

	if err != nil {
		return nil, err
	}

	q = bson.M{"_id": id}
	q2 = bson.M{"$addToSet": bson.M{"requestsSent": userID}}

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := collection.FindOneAndUpdate(ctx, q, q2, &opt)

	if result.Err() != nil {
		return nil, result.Err()
	}

	err = result.Decode(&results)

	if err != nil {
		return nil, err
	}

	return results, nil
}
