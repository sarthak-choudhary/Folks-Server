package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Event model
type Event struct {
	ID           primitive.ObjectID   `bson:"_id" json:"id" required:"true"`
	Name         string               `bson:"name" json:"name" required:"true"`
	Description  string               `bson:"description" json:"description,omitempty"`
	Destination  string               `bson:"destination" json:"destination,omitempty"`
	Datetime     time.Time            `bson:"datetime,omitempty" json:"datetime,omitempty"`
	CreatedOn    time.Time            `bson:"createdOn,omitempty" json:"createdOn,omitempty" required:"true"`
	HostedBy     primitive.ObjectID   `bson:"hostedBy" json:"hostedBy" required:"true"`
	Participants []primitive.ObjectID `bson:"participants,omitempty" json:"participants,omitempty"`
	PicturesUrls []string             `bson:"picturesUrls,omitempty" json:"picturesUrls,omitempty"`
}

//Events - Slice of event
type Events []Event
