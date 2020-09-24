package query

import (
	"context"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

//GetEvent - fetch event by _id
func GetEvent(_id string, client *mongo.Client) (interface{}, error) {
	var result models.Event
	var err error

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		return nil, err
	}

	q := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("events")
	err = collection.FindOne(ctx, q).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
