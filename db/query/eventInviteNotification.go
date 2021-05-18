package query

import (
	"context"
	"errors"
	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func EventInviteNotification(client *mongo.Client, senderId primitive.ObjectID, receiverId []primitive.ObjectID, event models.Event) (error, models.Notification) {
	var notif models.Notification
	notif.ID = primitive.NewObjectID()
	notif.Receiver	=	receiverId
	notif.Sender	=	senderId
	notif.Event		=	event.ID
	notif.Code		=	1
	loc, _ := time.LoadLocation("Asia/Kolkata")
	notif.NotificationTime = time.Now().In(loc)

	err, newNotif := SendNotification(notif, client)
	if err != nil {
		return err, models.Notification{}
	}
	if newNotif.ID == primitive.NilObjectID && newNotif.Code == 0 && newNotif.Sender == primitive.NewObjectID() && newNotif.Event == primitive.NewObjectID() && len(newNotif.Receiver) == 0 {
		return errors.New("Couldn't send notification"), models.Notification{}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("notification")
	_, err = collection.InsertOne(ctx, newNotif)
	if err != nil {
		return err, models.Notification{}
	}
	return nil, newNotif
}