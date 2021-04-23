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

//DeclineInvite allows user-queries to decline the invite for a particular event-queries
func DeclineInvite(id primitive.ObjectID, eventID primitive.ObjectID, client *mongo.Client) (models.User, error) {
	var results models.User
	var err error
	emptyUserObject := models.User{}

	q := bson.M{"_id": id}
	q2 := bson.M{"$pull": bson.M{"invitesReceived": eventID}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("users")
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := collection.FindOneAndUpdate(ctx, q, q2, &opt)

	if result.Err() != nil {
		return emptyUserObject, result.Err()
	}

	err = result.Decode(&results)

	if err != nil {
		return emptyUserObject, err
	}

	collection = client.Database("folks").Collection("events")
	q = bson.M{"_id": eventID}
	q2 = bson.M{"$pull": bson.M{"invitelist": id}}

	err = collection.FindOneAndUpdate(ctx, q, q2).Err()

	if err != nil {
		return emptyUserObject, err
	}

	return results, nil
}
