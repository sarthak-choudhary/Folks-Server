package notification_queries

import (
	"context"
	"errors"
	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func EventJoinNotification(client *mongo.Client, senderId primitive.ObjectID, receiverId []primitive.ObjectID, event models.Event) (error, models.Notification) {
	var notif models.Notification
	notif.ID = primitive.NewObjectID()
	notif.Receiver	=	receiverId
	notif.Sender	=	senderId
	notif.Event		=	event.ID
	notif.Code		=	2

	err, newNotif := notif.SendNotification(client)
	if err != nil {
		return err, models.Notification{}
	}
	if newNotif.ID == primitive.NilObjectID && newNotif.Code == 0 && newNotif.Sender == primitive.NewObjectID() && newNotif.Event == primitive.NewObjectID() && len(newNotif.Receiver) == 0 {
		return errors.New("Couldn't send notification"), models.Notification{}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("notification-queries")
	_, err = collection.InsertOne(ctx, newNotif)
	if err != nil {
		return err, models.Notification{}
	}
	return nil, newNotif
}
