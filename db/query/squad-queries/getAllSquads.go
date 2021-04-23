package squad_queries

import (
	"context"
	"time"

	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllSquads - fetches all squads feom db.
func GetAllSquads(client *mongo.Client) (models.Squads, error) {
	var result models.Squads
	var err error
	emptySquadsObject := models.Squads{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("squads")
	cur, err := collection.Find(ctx, bson.D{}, options.Find())

	if err != nil {
		return emptySquadsObject, err
	}

	for cur.Next(ctx) {
		var squad models.Squad
		err = cur.Decode(&squad)

		if err != nil {
			return emptySquadsObject, err
		}

		result = append(result, squad)
	}

	if err = cur.Err(); err != nil {
		return emptySquadsObject, err
	}

	cur.Close(ctx)
	return result, nil
}
