package query

import (
	"context"
	//"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// UpdateUser updates the user
func UpdateUser(u *models.User, client *mongo.Client) (*models.User, error) {
	var err error
	var results models.User
	emptyUserObject := models.User{}

	if err != nil {
		return &emptyUserObject, err
	}

	q := bson.M{"_id": u.ID}
	q2 := bson.M{"$set": bson.M{
		"firstName":       u.FirstName,
		"lastName":        u.LastName,
		"phoneNo":         u.PhoneNo,
		"interests":       u.Interests,
		"isComplete":      u.IsComplete,
		"followedByCount": u.FollowedByCount,
		"following":       u.Following,
		"events":          u.Events,
		"pictureUrls":     u.PicturesUrls,
		"isPublic":		   u.IsPublic,
		"username":		   u.Username,
	}}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("folks").Collection("users")
	err = collection.FindOneAndUpdate(ctx, q, q2, &opt).Decode(&results)

	if err != nil {
		return &emptyUserObject, err
	}

	return &results, nil
}

