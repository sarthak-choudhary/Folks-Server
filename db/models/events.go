package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Event model
type Event struct {
	ID                primitive.ObjectID   `bson:"_id" json:"_id" required:"true"`
	Name              string               `bson:"name" json:"name" required:"true"`
	Description       string               `bson:"description,omitempty" json:"description,omitempty"`
	Destination       string               `bson:"destination,omitempty" json:"destination,omitempty"`
	LocationLatitude  float32              `bson:"locationLatitude,omitempty" json:"locationLatitude,omitempty"`
	LocationLongitude float32              `bson:"locationLongitude,omitempty" json:"locationLongitude,omitempty"`
	Datetime          time.Time            `bson:"datetime,omitempty" json:"datetime,omitempty"`
	CreatedOn         time.Time            `bson:"createdOn" json:"createdOn" required:"true"`
	HostedBy          primitive.ObjectID   `bson:"hostedBy" json:"hostedBy" required:"true"`
	Participants      []primitive.ObjectID `bson:"participants,omitempty" json:"participants,omitempty"`
	PicturesUrls      []string             `bson:"picturesUrls,omitempty" json:"picturesUrls,omitempty"`
}

//Events - Slice of event
type Events []Event
