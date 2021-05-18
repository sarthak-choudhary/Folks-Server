package query

import (
	"context"
	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func MyNotifications(client *mongo.Client, user models.User)	(error, models.Notifications)	{
	var result models.Notifications
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("folks").Collection("notification")
	query := bson.M{
		"receiver": user.ID,
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"notificationTime", -1}})
	cur, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		return err, models.Notifications{}
	}
	for cur.Next(context.TODO())	{
		var item models.Notification
		err := cur.Decode(&item)
		if err != nil {
			return err, models.Notifications{}
		}
		result = append(result, item)
	}
	if err = cur.Err(); err != nil {
		return err, models.Notifications{}
	}
	cur.Close(context.TODO())
	return nil, result
}
