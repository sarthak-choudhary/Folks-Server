package models

import (
	e "github.com/wefolks/backend/db/query/event-queries"
	"github.com/wefolks/backend/db/query/user-queries"
	"github.com/wefolks/backend/fcm"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type Notification struct {
	ID				primitive.ObjectID			`json:"_id, omitempty" bson:"_id, omitempty"`
	Code			int64						`json:"code" bson:"code"`
	Sender			primitive.ObjectID			`json:"sender" bson:"sender"`
	Receiver		[]primitive.ObjectID		`json:"receiver" bson:"receiver"`
	Event			primitive.ObjectID			`json:"event-queries,omitempty" bson:"event-queries,omitempty"`
}

type Notifications []Notification

var MessageCode = map[int64]string	{
	1	:	"$sender$ has invited you to join $event_name$",
	2	:	"$sender$ wants to join your event-queries, $event_name$",
	3	:	"$sender$ has requested to follow you on the folks app",
	4	:	"$sender$ has accepted your follow request",
}

func (n Notification) SendNotification(client *mongo.Client) (error, Notification) {
	fcmClient := fcm.GetFcmClient()
	switch n.Code {
		case 1,2:	{
			sender, err := user_queries.GetUserByID(n.Sender.String(), client)
			if err != nil {
				return err, Notification{}
			}
			err, receivers  := user_queries.GetListOfUsersById(client, n.Receiver)
			if err != nil {
				return err, Notification{}
			}
			event, err := e.GetEvent(n.Event.String(), client)
			for _, r := range receivers {
				msg := MessageCode[n.Code]
				msg = strings.Replace(msg, "$sender$", sender.FirstName+" "+sender.LastName, 1)
				msg = strings.Replace(msg, "$event_name$", event.Name, 1)
				data := map[string]interface{}{
					"message": msg,
					"details": map[string]primitive.ObjectID{
					"sender":   n.Sender, "receiver": r.ID,
					"event-queries":    event.ID,
					},
				}
				_, err := fcm.SendMultipleNotifications(fcmClient, data, r.FcmToken)
				if err != nil {
					return err, Notification{}
				}
			}
		}
		case 3,4:	{
			sender, err := user_queries.GetUserByID(n.Sender.String(), client)
			if err != nil {
				return err, Notification{}
			}
			// Follow requests should only have a single receiver
			receiver, err := user_queries.GetUserByID(n.Receiver[0].String(), client)
			if err != nil {
				return err, Notification{}
			}
			msg := MessageCode[n.Code]
			msg = strings.Replace(msg, "$sender$", sender.FirstName, 1)
			data := map[string]interface{}{
				"message": msg,
				"details": map[string]primitive.ObjectID {
					"sender": n.Sender,
					"receiver": receiver.ID,
				},
			}
			_, err = fcm.SendMultipleNotifications(fcmClient, data, receiver.FcmToken)
			if err != nil{
				return err, Notification{}
			}
		}
		default:
			return nil, Notification{}
	}
	return nil, Notification{}
}