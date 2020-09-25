package query

import (
	"context"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//AddSquad - adds new squad to db.
func AddSquad(name string, description string, groupImages []string, admins []primitive.ObjectID, members []primitive.ObjectID, invitesSent []primitive.ObjectID, client *mongo.Client) (interface{}, error) {
	var err error
	var squad models.Squad

	squad.ID = primitive.NewObjectID()
	squad.Name = name
	squad.Description = description
	squad.GroupImages = groupImages
	squad.Admins = admins
	squad.Members = members
	squad.InvitesSent = invitesSent

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("squads")
	_, err = collection.InsertOne(ctx, squad)

	if err != nil {
		return nil, err
	}

	return squad, nil
}
