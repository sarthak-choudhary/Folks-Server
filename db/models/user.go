package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the User entity stored in the database
type User struct {
	ID               primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName        string               `json:"firstName" bson:"firstName"`
	LastName         string               `json:"lastName" bson:"lastName"`
	Email            string               `json:"email" bson:"email"`
	Password         string               `json:"password,omitempty" bson:"password,omitempty"`
	PhoneNo          string               `json:"phoneNo,omitempty" bson:"phoneNo,omitempty"`
	Bio              string               `json:"bio,omitempty" bson:"bio,omitempty"`
	Age              int64				  `json:"age,omitempty" bson:"age,omitempty"`
	Gender			 int64				  `json:"gender,omitempty" bson:"gender,omitempty"`
	Interests        []string             `json:"interests,omitempty" bson:"interests,omitempty"`
	IsComplete       bool                 `json:"isComplete,omitempty" bson:"isComplete,omitempty"`
	FollowedByCount  int64                `json:"followedByCount,omitempty" bson:"followedByCount,omitempty"`
	Following        []primitive.ObjectID `json:"following,omitempty" bson:"following,omitempty"`
	Events           []primitive.ObjectID `json:"events,omitempty" bson:"events,omitempty"`
	RequestsSent     []primitive.ObjectID `json:"requestsSent,omitempty" bson:"requestsSent,omitempty"`
	RequestsReceived []primitive.ObjectID `json:"requestsReceived,omitempty" bson:"requestsReceived,omitempty"`
	InvitesSent      []primitive.ObjectID `json:"invitesSent,omitempty" bson:"invitesSent,omitempty"`
	InvitesReceived  []primitive.ObjectID `json:"invitesReceived,omitempty" bson:"invitesReceived,omitempty"`
	PicturesUrls     []string             `bson:"picturesUrls,omitempty" json:"picturesUrls,omitempty"`
}

//Users - Slice of Users
type Users []User

//Claims used for jwt token
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

//GenerateJWT generates a jwt for a user
func (u User) GenerateJWT() (string, error) {
	expirationTime := time.Now().Add(90 * 24 * time.Hour)
	claims := &Claims{
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwtKey := []byte("wefolks12345")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
