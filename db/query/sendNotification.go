package query

import (
	"github.com/wefolks/backend/db/models"
	//user_queries "github.com/wefolks/backend/db/query/user"
	"github.com/wefolks/backend/fcm"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

func SendNotification(n models.Notification, client *mongo.Client) (error, models.Notification) {
	fcmClient := fcm.GetFcmClient()
	switch n.Code {
	case 1,2:	{
		sender, err := GetUserByID(n.Sender.String(), client)
		if err != nil {
			return err, models.Notification{}
		}
		err, receivers  := GetListOfUsersById(client, n.Receiver)
		if err != nil{
			return err, models.Notification{}
		}
		event, err := GetEvent(n.Event.String(), client)
		for _, r := range receivers {
			msg := models.MessageCode[n.Code]
			msg = strings.Replace(msg, "$sender$", sender.FirstName+" "+sender.LastName, 1)
			msg = strings.Replace(msg, "$event_name$", event.Name, 1)
			data := map[string]interface{}{
				"message": msg,
				"details": map[string]primitive.ObjectID{
					"sender":   n.Sender,
					"receiver": r.ID,
					"event":    event.ID,
				},
			}
			_, err := fcm.SendMultipleNotifications(fcmClient, data, r.FcmToken)
			if err != nil {
				return err, models.Notification{}
			}
		}
	}
	case 3,4:	{
		sender, err := GetUserByID(n.Sender.String(), client)
		if err != nil {
			return err, models.Notification{}
		}
		// Follow requests should only have a single receiver
		receiver, err := GetUserByID(n.Receiver[0].String(), client)
		if err != nil {
			return err, models.Notification{}
		}
		msg := models.MessageCode[n.Code]
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
			return err, models.Notification{}
		}
	}
	default:
		return nil, models.Notification{}
	}
	return nil, models.Notification{}
}
