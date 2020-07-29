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
	// Change mongo ApplyURI -> "mongodb://db:27017/events_microservices" to run with docker
	// Change mongo ApplyURI -> "mongodb://localhost:27017/events_microservices" to run locally
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://db:27017/events_microservice"))
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
		Events:  client.Database("events_microservice").Collection("events"),
	}
}

// CloseDB closes connection to the database
func (db MongoDB) CloseDB() {
	err := db.Session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
