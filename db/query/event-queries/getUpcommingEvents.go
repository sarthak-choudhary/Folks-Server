package event_queries

import (
	"context"
	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func GetUpcommingEvents(_id string,c *mongo.Client) (models.Events, error)	{
	var result models.Events
	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil	{
		return nil, err
	}

	q := bson.M{
		"participants": id,
		"datetime": bson.M{
			"$gt": time.Now(),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := c.Database("folks").Collection("events")

	res, err := collection.Find(ctx, q)
	if err!=nil{
		return nil, err
	}

	for res.Next(ctx)	{
		var event models.Event
		err = res.Decode(&event)
		if err != nil {
			return nil, err
		}
		result = append(result, event)
	}
	return  result, nil
}
