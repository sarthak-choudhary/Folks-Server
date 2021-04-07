package query

import (
	"context"
	"fmt"
	"github.com/WeFolks/search_service/grpc"
	g "google.golang.org/grpc"
	"log"
	"time"

	"github.com/anshalshukla/folks/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//AddEvent - adds new event to db.
func AddEvent(name string, description string, destination string, locationLatitude float64, locationLongitude float64, datetime time.Time, hostedBy primitive.ObjectID, inviteList []primitive.ObjectID, picturesUrls []string, admins []primitive.ObjectID, owner string, client *mongo.Client) (models.Event, error) {
	var err error
	var event models.Event
	emptyEventObject := models.Event{}
	event.ID = primitive.NewObjectID()
	event.Name = name
	event.Description = description
	event.Destination = destination
	event.LocationLatitude = locationLatitude
	event.LocationLongitude = locationLongitude
	event.Datetime = datetime
	event.CreatedOn = time.Now()
	event.HostedBy = hostedBy
	event.InviteList = inviteList
	event.Admins = admins
	event.Participants=admins
	event.PicturesUrls = picturesUrls

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database("folks").Collection("events")
	_, err = collection.InsertOne(ctx, event)
	if err != nil {
		return emptyEventObject, err
	}

	var conn *g.ClientConn
	conn, err = g.Dial("3.142.74.30:9000", g.WithInsecure())

	if err != nil {
		fmt.Print("Connection not established\n")
		log.Fatalf("Object could not be added in search Database\n")
		return event, err
	}

	defer conn.Close()

	c := grpc.NewSearchServiceClient(conn)

	item := grpc.Item{
		Id:          event.ID.Hex(),
		Name:        event.Name,
		Owner:       owner,
		Description: event.Description,
		Type:        1,
	}

	_, err = c.AddItem(context.Background(), &item)

	if err != nil {
		fmt.Print("This is the problem\n")
		log.Fatalf("Object could not be added in search Database\n")
		return event, err
	}

	return event, nil
}
