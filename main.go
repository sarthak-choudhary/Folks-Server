package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anshalshukla/folks/api"
	"github.com/anshalshukla/folks/db"
	"github.com/anshalshukla/folks/gql"
	"github.com/anshalshukla/folks/middleware"
	"github.com/anshalshukla/folks/pkg/elasticsearch"
	"github.com/anshalshukla/folks/util"
	"github.com/graphql-go/handler"
)

var dbConnection db.MongoDB

func main() {
	// MongoDB client generation
	dbConnection = db.ConnectDB()
	defer dbConnection.CloseDB()
	client := dbConnection.Session

	// Elasticsearch client generation
	elastiClient, err := elasticsearch.GetESClient()
	if err != nil {
		log.Fatal("Elastic Search client can't be setup", err)
		return
	}

	// Graphql handler setup
	schema := gql.InitSchema(&dbConnection, elastiClient)
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Test case for gcp load balancer
	http.Handle("/test", middleware.LogReq(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "API Live")
	})))

	http.Handle("/graphql", middleware.LogReq(middleware.Auth(client, h)))
	http.Handle("/graphql_test", h)
	http.Handle("/sign_up", middleware.LogReq(api.SignUp(client, elastiClient)))
	http.Handle("/login", middleware.LogReq(api.Login(client)))
	http.Handle("/google_login", middleware.LogReq(api.GoogleOauth(client, elastiClient)))
	http.Handle("/search", middleware.LogReq(middleware.Auth(client, api.GetData(elastiClient))))
	http.Handle("/my_profile", middleware.LogReq(middleware.Auth(client, api.Myprofile())))
	http.HandleFunc("/image_upload", util.Handler)

	log.Printf("HTTP server started on :%s", os.Getenv("PORT_FOR_WEBAPP"))
	addr := ":" + string(os.Getenv("PORT_FOR_WEBAPP"))
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
