package query

import (
	"context"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func AcceptRequest(id primitive.ObjectID, user_id string, client *mongo.Client) (interface{}, error) {
	var err error
	var results models.User

	userID, err := primitive.ObjectIDFromHex(user_id)

	if err != nil {
		return nil, err
	}

	q := bson.M{"_id": id}
	q2 := bson.M{"$pull": bson.M{"requestsReceived": userID}, "$inc": bson.M{"followedByNumber": 1}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("users")
	err = collection.FindOneAndUpdate(ctx, q, q2).Decode(&results)

	q = bson.M{"_id": userID}
	q2 = bson.M{"$addToSet": bson.M{"following": id}, "$pull": bson.M{"requestsSent": id}}
	err = collection.FindOneAndUpdate(ctx, q, q2).Err()

	return results, nil
}
