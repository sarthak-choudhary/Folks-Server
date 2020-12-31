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

func AcceptRequest(id primitive.ObjectID, userID primitive.ObjectID, client *mongo.Client) (interface{}, error) {
	var results models.User
	var err error

	q := bson.M{"_id": id}
	q2 := bson.M{"$pull": bson.M{"requestsReceived": userID}, "$inc": bson.M{"followedByCount": 1}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("users")

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

	q = bson.M{"_id": userID}
	q2 = bson.M{"$addToSet": bson.M{"following": id}, "$pull": bson.M{"requestsSent": id}}
	err = collection.FindOneAndUpdate(ctx, q, q2).Err()

	if err != nil {
		return nil, err
	}

	return results, nil
}
