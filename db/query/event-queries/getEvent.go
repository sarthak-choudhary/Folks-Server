package event_queries

import (
	"context"
	"time"

	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

//GetEvent - fetch event-queries by _id
func GetEvent(_id string, client *mongo.Client) (models.Event, error) {
	var result models.Event
	var err error
	emptyEventObject := models.Event{}

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		return emptyEventObject, err
	}

	q := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("events")
	err = collection.FindOne(ctx, q).Decode(&result)

	if err != nil {
		return emptyEventObject, err
	}

	return result, nil
}
