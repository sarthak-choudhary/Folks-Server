package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/anshalshukla/folks/db/models"
	"github.com/anshalshukla/folks/db/query"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)

//Auth checks for the Authorization header, decodes
//the token and adds the user to request context.
func Auth(client *mongo.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		user, _ := query.GetUser(claims.Email, client)
		key := "user"
		ctx := context.WithValue(r.Context(), key, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
