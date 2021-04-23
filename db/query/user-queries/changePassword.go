package user_queries

import (
	"context"
	"errors"
	"fmt"
	"github.com/wefolks/backend/db/models"
	"github.com/wefolks/backend/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// changePassword: Change password for authenticated user-queries
// np -> New password
// op -> Old password
func ChangePassword(op string, np string, _id primitive.ObjectID, c *mongo.Client)	(models.User, error)	{
	// Get user-queries from id
	fmt.Print("I was here")
	stringUserID := _id.Hex()
	user, err := GetUserByID(stringUserID, c)
	fmt.Print("I was here***")
	if err!=nil {
		return models.User{}, err
	}
	fmt.Print("I was here****")
	if !util.MatchesWithHash(op, user.Password) {
		return models.User{}, errors.New("Incorrect Password")
	}
	fmt.Print("I was here-2")
	userCollection := c.Database("folks").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	passwordHash, _ := util.HashPassword(np)
	fmt.Print("I was here-3")
	err = userCollection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": _id},
		bson.D{
			{"$set", bson.D{{"password", passwordHash}}},
		},
	).Decode(&user)
	fmt.Print("I was here-4")
	if err != nil{
		return models.User{}, err
	}

	return user, nil

}