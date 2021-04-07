package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anshalshukla/folks/db/models"
	"github.com/anshalshukla/folks/db/query"
	"github.com/anshalshukla/folks/util"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/oauth2/v1"
)

// SignUp endpoint to add user to the database
func SignUp(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			payload := struct {
				Error string `json:"error"`
			}{Error: "Invalid or Incomplete fields"}

			json.NewEncoder(w).Encode(payload)
			return
		}

		_, err = query.GetUser(user.Email, client)

		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			payload := struct {
				Error string `json:"error"`
			}{Error: "User with this email already exists"}

			json.NewEncoder(w).Encode(payload)
			return
		}

		_, err = query.GetUserByPhoneNo(user.PhoneNo, client)

		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			payload := struct {
				Error string `json:"error"`
			}{Error: "User with this phone number already exists"}

			json.NewEncoder(w).Encode(payload)
			return
		}

		_, err = query.GetUserByUsername(user.Username, client)

		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			payload := struct {
				Error string `json:"error"`
			}{Error: "User with this username already exists"}

			json.NewEncoder(w).Encode(payload)
			return
		}

		user.Password, _ = util.HashPassword(user.Password)
		user.IsComplete = true
		user.IsPublic	=	false

		_id, err := query.AddUser(&user, client)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		token, err := user.GenerateJWT()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		payload := struct {
			ID        string   `json:"_id,omitempty" bson:"_id,omitempty"`
			FirstName string   `json:"firstName"`
			LastName  string   `json:"lastName"`
			Email     string   `json:"email"`
			PhoneNo   string   `json:"phoneNo"`
			Bio       string   `json:"bio"`
			Gender    int64    `json:"gender"`
			Age		  int64    `json:"age"`
			Interests []string `json:"interests"`
			Token     string   `json:"token"`
			Username  string   `json:"username"`
		}{
			ID:        _id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			PhoneNo:   user.PhoneNo,
			Bio:       user.Bio,
			Gender:    user.Gender,
			Age:       user.Age,
			Interests: user.Interests,
			Token:     token,
			Username:  user.Username,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(payload)

	})
}

// Login endpoint to check user in the database
func Login(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := query.GetUser(data.Email, client)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			payload := struct {
				Error string `json:"error"`
			}{Error: "Invalid email or password"}

			json.NewEncoder(w).Encode(payload)
			return
		}

		if !util.MatchesWithHash(data.Password, user.Password) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			payload := struct {
				Error string `json:"error"`
			}{Error: "Invalid email or password"}

			json.NewEncoder(w).Encode(payload)
			return
		}

		token, err := user.GenerateJWT()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		payload := struct {
			ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
			FirstName string   `json:"firstName"`
			LastName  string   `json:"lastName"`
			Email     string   `json:"email"`
			PhoneNo   string   `json:"phoneNo"`
			Bio       string   `json:"bio"`
			Gender    int64   `json:"gender"`
			Age		  int64   `json:"age"`
			Interests []string `json:"interests"`
			Token     string   `json:"token"`
		}{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			PhoneNo:   user.PhoneNo,
			Bio:       user.Bio,
			Gender:    user.Gender,
			Age:       user.Age,
			Interests: user.Interests,
			Token:     token,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(payload)

	})
}

//GoogleOauth enables signup with google
func GoogleOauth(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			IDToken string `json:"id_token"`
		}

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var httpClient = &http.Client{}
		oauth2Service, err := oauth2.New(httpClient)
		tokenInfoCall := oauth2Service.Tokeninfo()
		tokenInfoCall.IdToken(data.IDToken)
		_, err = tokenInfoCall.Do()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		type TokenInfo struct {
			Iss string `json:"iss"`
			// userId
			Sub string `json:"sub"`
			Azp string `json:"azp"`
			// clientId
			Aud string `json:"aud"`
			Iat int64  `json:"iat"`
			// expired time
			Exp int64 `json:"exp"`

			Email         string `json:"email"`
			EmailVerified bool   `json:"email_verified"`
			AtHash        string `json:"at_hash"`
			Name          string `json:"name"`
			GivenName     string `json:"given_name"`
			FamilyName    string `json:"family_name"`
			Picture       string `json:"picture"`
			Local         string `json:"locale"`
			jwt.StandardClaims
		}

		token, _, err := new(jwt.Parser).ParseUnverified(data.IDToken, &TokenInfo{})
		tokenInfo, ok := token.Claims.(*TokenInfo)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		var user *models.User
		var jwtToken string
		var _id primitive.ObjectID

		user, err = query.GetUser(tokenInfo.Email, client)

		if user != nil {
			_id = user.ID
			jwtToken, err = user.GenerateJWT()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			payload := struct {
				ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
				FirstName  string             `json:"firstName"`
				LastName   string             `json:"lastName"`
				Email      string             `json:"email"`
				Token      string             `json:"token"`
				IsComplete bool               `json:"isComplete"`
			}{
				ID:         _id,
				FirstName:  user.FirstName,
				LastName:   user.LastName,
				Email:      user.Email,
				Token:      jwtToken,
				IsComplete: user.IsComplete,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(payload)

		} else {
			var newUser models.User

			newUser.Email = tokenInfo.Email
			newUser.FirstName = tokenInfo.GivenName
			newUser.LastName = tokenInfo.FamilyName
			newUser.IsComplete = false

			_id, err := query.AddUser(&newUser, client)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			jwtToken, err = newUser.GenerateJWT()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			payload := struct {
				ID         string `json:"_id,omitempty" bson:"_id,omitempty"`
				FirstName  string `json:"firstName"`
				LastName   string `json:"lastName"`
				Email      string `json:"email"`
				Token      string `json:"token"`
				IsComplete bool   `json:"isComplete"`
			}{
				ID:         _id,
				FirstName:  newUser.FirstName,
				LastName:   newUser.LastName,
				Email:      newUser.Email,
				Token:      jwtToken,
				IsComplete: newUser.IsComplete,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(payload)
		}
	})
}

//Myprofile function gives you the profile of logged in user
func Myprofile() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)

		payload := struct {
			ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
			FirstName string             `json:"firstName"`
			LastName  string             `json:"lastName"`
			Email     string             `json:"email"`
			PhoneNo   string             `json:"phoneNo"`
			Interests []string           `json:"interests"`
		}{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			PhoneNo:   user.PhoneNo,
			Interests: user.Interests,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(payload)
	})
}
