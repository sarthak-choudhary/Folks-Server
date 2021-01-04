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

//AssignAdmin function allows an admin to give admin status to other users also.
func AssignAdmin(eventID primitive.ObjectID, admin_id []primitive.ObjectID, client *mongo.Client) (models.Event, error) {
	var results models.Event
	var err error
	emptyEventObject := models.Event{}

	q := bson.M{"_id": eventID}
	q2 := bson.M{"$addToSet": bson.M{"admins": bson.M{"$each": admin_id}}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("events")
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

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
