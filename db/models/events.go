package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Event model
type Event struct {
	ID                primitive.ObjectID   `bson:"_id" json:"_id" required:"true"`
	Category          int64                `bson:"category,omitempty" json:"category,omitempty"` // 0 - for open, 1 - for closed, 2 - for invite-only
	Name              string               `bson:"name" json:"name" required:"true"`
	Description       string               `bson:"description,omitempty" json:"description,omitempty"`
	Destination       string               `bson:"destination,omitempty" json:"destination,omitempty"`
	LocationLatitude  float64              `bson:"locationLatitude,omitempty" json:"locationLatitude,omitempty"`
	LocationLongitude float64              `bson:"locationLongitude,omitempty" json:"locationLongitude,omitempty"`
	Datetime          time.Time            `bson:"datetime,omitempty" json:"datetime,omitempty"`
	CreatedOn         time.Time            `bson:"createdOn" json:"createdOn" required:"true"`
	HostedBy          primitive.ObjectID   `bson:"hostedBy" json:"hostedBy" required:"true"`
	Admins            []primitive.ObjectID `bson:"admins" json:"admins" required:"true"`
	Participants      []primitive.ObjectID `bson:"participants,omitempty" json:"participants,omitempty"`
	Waitlist          []primitive.ObjectID `bson:"waitlist,omitempty" json:"waitlist,omitempty"`
	InviteList        []primitive.ObjectID `bson:"invitelist,omitempty" json:"invitelist,omitempty"`
	PicturesUrls      []string             `bson:"picturesUrls,omitempty" json:"picturesUrls,omitempty"`
}

//Events - Slice of event
type Events []Event
