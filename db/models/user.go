package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the User entity stored in the database
type User struct {
	ID              primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName       string               `json:"firstName"`
	LastName        string               `json:"lastName"`
	Email           string               `json:"email"`
	Password        string               `json:"password,omitempty"`
	PhoneNo         string               `json:"phoneNo,omitempty"`
	Interests       []string             `json:"interests,omitempty"`
	IsComplete      bool                 `json:"isComplete,omitempty"`
	FollowedByCount int64                `json:"followedByCount,omitempty"`
	Following       []primitive.ObjectID `json:"following,omitempty"`
	Events          []primitive.ObjectID `json:"events,omitempty"`
	PicturesUrls    []string             `bson:"picturesUrls,omitempty" json:"picturesUrls,omitempty"`
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
