package query

import (
	"context"
	"errors"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// UpdateUser updates the user
func UpdateUser(u *models.User, client *mongo.Client) (models.User, error) {
	var err error
	var results models.User
	emptyUserObject := models.User{}

	if err != nil {
		return emptyUserObject, err
	}

	q := bson.M{"_id": u.ID}
	q2 := bson.M{"$set": bson.M{
		"firstname":       u.FirstName,
		"lastname":        u.LastName,
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("folks").Collection("users")
	err = collection.FindOneAndUpdate(ctx, q, q2).Decode(&results)

	if err != nil {
		return emptyUserObject, err
	}

	err = errors.New("User info can only be modified only by user himself")

	results.FirstName = u.FirstName
	results.LastName = u.LastName
	results.PhoneNo = u.PhoneNo
	results.Interests = u.Interests
	results.IsComplete = u.IsComplete
	results.FollowedByCount = u.FollowedByCount
	results.Following = u.Following
	results.Events = u.Events
	results.PicturesUrls = u.PicturesUrls
	results.Username = u.Username
	results.IsPublic = u.IsPublic

	return results, nil
}

