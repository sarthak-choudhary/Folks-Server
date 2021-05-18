package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chatroom struct {
	ID				primitive.ObjectID					`json:"_id,omitempty" bson:"_id,omitempty"`
	IsGroup			bool								`json:"isGroup,required" bson:"isGroup,required"`
	Members			[]primitive.ObjectID				`json:"members,required" bson:"members,required"`
}
