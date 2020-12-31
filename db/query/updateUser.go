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
func UpdateUser(_id string, firstname string, lastname string, phoneNo string, interests []string, isComplete bool, followedByCount int64, following []primitive.ObjectID, events []primitive.ObjectID, picturesUrls []string, client *mongo.Client) (interface{}, error) {
	var err error
	var results models.User

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id}

	q2 := bson.M{"$set": bson.M{
		"firstname":       firstname,
		"lastname":        lastname,
		"phoneNo":         phoneNo,
		"interests":       interests,
		"isComplete":      isComplete,
		"followedByCount": followedByCount,
		"following":       following,
		"events":          events,
		"pictureUrls":     picturesUrls,
	}}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("folks").Collection("users")
	err = collection.FindOneAndUpdate(ctx, q, q2).Decode(&results)
	if err != nil {
		return nil, err
	}

	err = errors.New("User info can only be modified only by user himself.")

	results.FirstName = firstname
	results.LastName = lastname
	results.PhoneNo = phoneNo
	results.Interests = interests
	results.IsComplete = isComplete
	results.FollowedByCount = followedByCount
	results.Following = following
	results.Events = events
	results.PicturesUrls = picturesUrls
	return results, nil
}
