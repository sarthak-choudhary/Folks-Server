package notification_queries

import (
	"context"
	"errors"
	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func FollowRequestNotification(client *mongo.Client, receiverId primitive.ObjectID, senderId primitive.ObjectID) (error, models.Notification){
	var notif models.Notification
	notif.Code = 3
	notif.Sender = senderId
	notif.Receiver = append(notif.Receiver, receiverId)
	err, newNotif := notif.SendNotification(client)
	if err != nil	{
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