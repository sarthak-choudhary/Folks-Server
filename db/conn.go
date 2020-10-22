package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB holds session and database info.
type MongoDB struct {
	Session *mongo.Client
	Events  *mongo.Collection
}

// ConnectDB connects to the database
func ConnectDB() MongoDB {
	// Change mongo ApplyURI -> "mongodb://db:27017/folks" to run with docker
	// Change mongo ApplyURI -> "mongodb://localhost:27017/folks" to run locally
	// Change mongo ApplyURI -> "mongodb+srv://folks:folks@cluster0.9up09.mongodb.net/<dbname>?retryWrites=true&w=majority" to run on mongodb atlas
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://folks:folks@cluster0.9up09.mongodb.net/<dbname>?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return MongoDB{
		Session: client,
	}
}

// CloseDB closes connection to the database
func (db MongoDB) CloseDB() {
	err := db.Session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
