package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notification struct {
	ID					primitive.ObjectID			`json:"_id, omitempty" bson:"_id, omitempty"`
	Code				int64						`json:"code" bson:"code"`
	Sender				primitive.ObjectID			`json:"sender" bson:"sender"`
	Receiver			[]primitive.ObjectID		`json:"receiver" bson:"receiver"`
	Event				primitive.ObjectID			`json:"event,omitempty" bson:"event,omitempty"`
	NotificationTime	time.Time					`json:"notificationTime,required" bson:"notificationTime,required"`
}

type Notifications []Notification

var MessageCode = map[int64]string	{
	1	:	"$sender$ has invited you to join $event_name$",
	2	:	"$sender$ wants to join your event, $event_name$",
	3	:	"$sender$ has requested to follow you on the folks app",
	4	:	"$sender$ has accepted your follow request",
}

