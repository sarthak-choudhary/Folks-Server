package query

import (
	"context"
	"time"

	"github.com/wefolks/backend/db/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllEvents - fetches all events from db.
func GetAllEvents(client *mongo.Client) (models.Events, error) {
	var result models.Events
	var err error
	emptyEventsObject := models.Events{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("events")
	cur, err := collection.Find(ctx, bson.D{}, options.Find())

	if err != nil {
		return emptyEventsObject, err
	}

	for cur.Next(ctx) {
		var event models.Event
		err = cur.Decode(&event)
		if err != nil {
			return nil, err
		}
		result = append(result, event)
	}

	if err = cur.Err(); err != nil {
		return emptyEventsObject, err
	}

	cur.Close(ctx)
	return result, nil
}
