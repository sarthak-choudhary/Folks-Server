package query

import (
	"context"
	"errors"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// UpdateEvent updates the event
func UpdateEvent (_id string, name string, description string, destination string, locationLatitude float32, locationLongitude float32, datetime time.Time, userID primitive.ObjectID, participants []primitive.ObjectID, picturesUrls []string, client *mongo.Client) (interface{}, error) {
	var err error
	var results models.Event

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id}

	q2 := bson.M{"$set": bson.M{
		"name":              name,
		"description":       description,
		"destination":       destination,
		"locationLatitude":  locationLatitude,
		"locationLongitude": locationLongitude,
		"datetime":          datetime,
		"participants":      participants,
		"pictureUrls":       picturesUrls,
	}}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("folks").Collection("events")
	err = collection.FindOneAndUpdate(ctx, q, q2).Decode(&results)
	if err != nil {
		return nil, err
	}

	err = errors.New("Event can only be modified by user who created it")

	results.Name = name
	results.Description = description
	results.Destination = destination
	results.LocationLatitude = locationLatitude
	results.LocationLongitude = locationLongitude
	results.Datetime = datetime
	results.PicturesUrls = picturesUrls
	return results, nil
}
