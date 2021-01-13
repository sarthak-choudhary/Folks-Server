package util

import (
	"context"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// NearbyEvents - returns all the events in the circle of radius(r) of latitude(lat) and longitude(lon)
func NearbyEvents(client *mongo.Client, lat float64, lon float64, r float64) (interface{}, error) {
	var result models.Events
	latMin := lat - r
	latMax := lat + r
	lonMin := lon - r
	lonMax := lon + r

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("events")
	cur, err := collection.Find(ctx, bson.M{"locationLatitude": bson.M{"$gt": latMin, "$lt": latMax}, "locationLongitude": bson.M{"$gt": lonMin, "$lt": lonMax}}, options.Find())

	if err != nil {
		return nil, err
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
		return nil, err
	}
	cur.Close(ctx)
	return result, nil
}

// NearbyEvents - returns all the events in the circle of radius(r) of latitude(lat) and longitude(lon)
func NearbyEventsWithImages(client *mongo.Client, lat float64, lon float64, r float64) (interface{}, error) {
	var result models.Events
	latMin := lat - r
	latMax := lat + r
	lonMin := lon - r
	lonMax := lon + r

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("events")
	cur, err := collection.Find(ctx, bson.M{"locationLatitude": bson.M{"$gt": latMin, "$lt": latMax}, "locationLongitude": bson.M{"$gt": lonMin, "$lt": lonMax}}, options.Find())

	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var event models.Event
		err = cur.Decode(&event)
		if err != nil {
			return nil, err
		}
		if len(event.PicturesUrls) > 0 {
			result = append(result, event)
		}
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return result, nil
}
