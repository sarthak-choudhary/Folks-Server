package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anshalshukla/events_mongodb/db"
)

var dbConnection db.MongoDB

func main() {
	dbConnection = db.ConnectDB()
	defer dbConnection.CloseDB()

	fmt.Println("Server started on port 8000!")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Listen And Serve:", err)
	}
}
