package main

import (
	"log"
	"net/http"

	"github.com/anshalshukla/folks/api"
	"github.com/anshalshukla/folks/db"
	"github.com/anshalshukla/folks/gql"
	"github.com/anshalshukla/folks/middleware"
	"github.com/graphql-go/handler"
)

var dbConnection db.MongoDB

func main() {
	dbConnection = db.ConnectDB()
	defer dbConnection.CloseDB()

	schema := gql.InitSchema(&dbConnection)
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	client := dbConnection.Session

	http.Handle("/graphql", h)
	http.Handle("/sign_up", middleware.LogReq(api.SignUp(client)))
	http.Handle("/login", middleware.LogReq(api.Login(client)))
	http.Handle("/google_login", middleware.LogReq(api.GoogleOauth(client)))
	http.Handle("/my_profile", middleware.LogReq(middleware.Auth(client, api.Myprofile())))

	log.Println("HTTP server started on :8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
