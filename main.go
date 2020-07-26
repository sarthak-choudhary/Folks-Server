package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anshalshukla/events_mongodb/db"
	"github.com/anshalshukla/events_mongodb/gql"
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

	http.Handle("/graphql", h)

	fmt.Println("Server started on port 8000!")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Listen And Serve:", err)
	}
}
