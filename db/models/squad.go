package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Squad model
type Squad struct {
	ID          primitive.ObjectID   `bson:"_id" json:"_id" required:"true"`
	Name        string               `bson:"name" json:"name" required:"true"`
	Description string               `bson:"description,omitempty" json:"description,omitempty"`
	GroupImages []string             `bson:"groupImages,omitempty" json:"groupImages,omitempty"`
	Admins      []primitive.ObjectID `bson:"admins" json:"admins" required:"true"`
	Members     []primitive.ObjectID `bson:"members,omitempty" json:"members,omitempty"`
	InvitesSent []primitive.ObjectID `bson:"invitesSent,omitempty" json:"invitesSent,omitempty"`
}

//Squads - Slice of squads
type Squads []Squad
