package event_queries

import (
	"context"
	"errors"
	notification_queries "github.com/wefolks/backend/db/query/notification-queries"
	"time"

	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

//InviteParticipants function allows admins to invite new participants for a particular event-queries
func InviteParticipants(eventID primitive.ObjectID, users []primitive.ObjectID, client *mongo.Client) (models.Event, error) {
	var results models.Event
	var err error
	emptyEventObject := models.Event{}

	q := bson.M{"_id": eventID}
	q2 := bson.M{"$addToSet": bson.M{"invitelist": bson.M{"$each": users}}}

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

	collection = client.Database("folks").Collection("users")

	for _, userID := range users {
		q = bson.M{"_id": userID}
		q2 = bson.M{"$addToSet": bson.M{"invitesReceived": eventID}}

		result = collection.FindOneAndUpdate(ctx, q, q2)

		if result.Err() != nil {
			return emptyEventObject, result.Err()
		}
	}
	event, err := GetEvent(eventID.String(), client)
	if err!=nil{
		return models.Event{}, err
	}
	err, _ = notification_queries.EventInviteNotification(client, event.HostedBy, users, event)
	if err != nil {
		return models.Event{}, errors.New("Error: Notification could not be sent.\n"+err.Error())
	}
	return results, nil
}
