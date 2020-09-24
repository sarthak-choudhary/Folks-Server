package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anshalshukla/folks/api"
	"github.com/anshalshukla/folks/db"
	"github.com/anshalshukla/folks/gql"
	"github.com/anshalshukla/folks/middleware"
	"github.com/anshalshukla/folks/util"
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
	// Test case for gcp load balancer
	http.Handle("/test", middleware.LogReq(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "API Live")
	})))

	http.Handle("/graphql", middleware.LogReq(middleware.Auth(client, h)))
	http.Handle("/sign_up", middleware.LogReq(api.SignUp(client)))
	http.Handle("/login", middleware.LogReq(api.Login(client)))
	http.Handle("/google_login", middleware.LogReq(api.GoogleOauth(client)))
	http.Handle("/my_profile", middleware.LogReq(middleware.Auth(client, api.Myprofile())))
	http.HandleFunc("/image_upload", util.Handler)

	log.Println("HTTP server started on :4000")
	err := http.ListenAndServe(":4000", nil)

	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
