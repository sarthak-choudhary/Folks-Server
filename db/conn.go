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
	session *mongo.Client
	events  *mongo.Collection
}

// ConnectDB connects to the database
func ConnectDB() MongoDB {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017/events_microservice"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return MongoDB{
		session: client,
		events:  client.Database("events_microservice").Collection("events"),
	}
}

// CloseDB closes connection to the database
func (db MongoDB) CloseDB() {
	err := db.session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
