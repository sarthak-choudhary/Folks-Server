package event_queries

import (
	"context"
	"time"

	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// UpdateEvent updates the event-queries
func UpdateEvent(id primitive.ObjectID, name string, description string, destination string, locationLatitude float64, locationLongitude float64, datetime time.Time, picturesUrls []string, client *mongo.Client) (models.Event, error) {
	var err error
	var results models.Event
	emptyEventObject := models.Event{}

	q := bson.M{"_id": id}
	q2 := bson.M{"$set": bson.M{
		"name":              name,
		"description":       description,
		"destination":       destination,
		"locationLatitude":  locationLatitude,
		"locationLongitude": locationLongitude,
		"datetime":          datetime,
		"pictureUrls":       picturesUrls,
	}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	collection := client.Database("folks").Collection("events")
	result := collection.FindOneAndUpdate(ctx, q, q2, &opt)

	if result.Err() != nil {
		return emptyEventObject, result.Err()
	}

	err = result.Decode(&results)

	if err != nil {
		return emptyEventObject, err
	}

	return results, nil
}
