package query

import (
	"context"
	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func GetListOfUsersById(client *mongo.Client, ids []primitive.ObjectID) (error, models.Users)	{
	var result models.Users

	query := bson.M{
		"_id" : bson.M{
			"$in": ids,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("users")
	cur, err := collection.Find(ctx, query)
	if err != nil {
		return err, models.Users{}
	}
	for cur.Next(ctx)	{
		var tempUser models.User
		cur.Decode(&tempUser)
		result = append(result, tempUser)
	}

	return nil, result
}
