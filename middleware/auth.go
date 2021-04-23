package middleware

import (
	"context"
	"encoding/json"
	"github.com/wefolks/backend/db/query/user-queries"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/wefolks/backend/db/models"
	"go.mongodb.org/mongo-driver/mongo"
)

//Auth checks for the Authorization header, decodes
//the token and adds the user-queries to request context.
func Auth(client *mongo.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		_, ok := headers["Authorization"]
		if !ok {
			payload := struct {
				Error string `json:"error"`
			}{Error: "Please login or sign up"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(payload)
			return
		}
		token := strings.Split(r.Header.Get("Authorization"), " ")[1]
		claims := &models.Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("wefolks12345"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, _ := user_queries.GetUser(claims.Email, client)
		key := "user-queries"
		ctx := context.WithValue(r.Context(), key, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
