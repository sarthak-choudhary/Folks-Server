package query

import (
	"context"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

//GetSquad - fetch squad by _id
func GetSquad(_id string, client *mongo.Client) (models.Squad, error) {
	var result models.Squad
	var err error
	emptySquadObject := models.Squad{}

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		return emptySquadObject, err
	}

	q := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("squads")
	err = collection.FindOne(ctx, q).Decode(&result)

	if err != nil {
		return emptySquadObject, err
	}

	return result, nil
}
