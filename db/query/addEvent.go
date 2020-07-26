package query

import (
	"context"
	"time"

	"github.com/anshalshukla/events_mongodb/db"
	"github.com/anshalshukla/events_mongodb/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddEvent - adds new event to db.
func AddEvent(name string, description string, destination string, locationLatitude float32, locationLongitude float32, datetime time.Time, hostedBy primitive.ObjectID, participants []primitive.ObjectID, picturesUrls []string, d *db.MongoDB) (interface{}, error) {
	var err error
	var event models.Event

	event.ID = primitive.NewObjectID()
	event.Name = name
	event.Description = description
	event.Destination = destination
	event.LocationLatitude = locationLatitude
	event.LocationLongitude = locationLongitude
	event.Datetime = datetime
	event.CreatedOn = time.Now()
	event.HostedBy = hostedBy
	event.Participants = participants
	event.PicturesUrls = picturesUrls

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err = d.Events.InsertOne(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}
