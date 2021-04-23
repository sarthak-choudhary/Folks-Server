package query

import (
	"context"
	"errors"
	"time"

	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

//AcceptRequest function accepts the follow request of particular user
func AcceptRequest(acceptorId primitive.ObjectID, requesterId primitive.ObjectID, client *mongo.Client) (models.User, error) {
	var results models.User
	var err error
	emptyUserObject := models.User{}

	q := bson.M{"_id": acceptorId}
	q2 := bson.M{"$pull": bson.M{"requestsReceived": requesterId}, "$inc": bson.M{"followedByCount": 1}}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("users")

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	result := collection.FindOneAndUpdate(ctx, q, q2, &opt)

	if result.Err() != nil {
		return emptyUserObject, result.Err()
	}

	err = result.Decode(&results)

	if err != nil {
		return emptyUserObject, err
	}

	q = bson.M{"_id": requesterId}
	q2 = bson.M{"$addToSet": bson.M{"following": acceptorId}, "$pull": bson.M{"requestsSent": acceptorId}}
	err = collection.FindOneAndUpdate(ctx, q, q2).Err()

	if err != nil {
		return emptyUserObject, err
	}

	err, _ = FollowAcceptedNotification(client, requesterId, acceptorId)
	if err != nil {
		return emptyUserObject, errors.New("Error: Notification could not be sent.\n"+err.Error())
	}
	
	return results, nil
}
