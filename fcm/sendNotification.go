package fcm

import (
	"errors"

	f "github.com/douglasmakey/go-fcm"
)

//SendSingleNotification - allows to send notification to a single device
func SendSingleNotification(client *f.Client, mssg interface{}, id string) (interface{}, error) {
	client.PushSingle(id, mssg)

	badRegistrations := client.CleanRegistrationIds()

	if len(badRegistrations) != 0 {
		err := errors.New("Invlaid FCM id(s)")
		return nil, err
	}

	status, err := client.Send()

	if err != nil {
		return nil, err
	}

	return status, err
}

//SendMultipleNotifications - allows to send notification to Multiple devices
func SendMultipleNotifications(client *f.Client, mssg interface{}, ids []string) (interface{}, error) {
	client.PushMultiple(ids, mssg)

	badRegistrations := client.CleanRegistrationIds()

	if len(badRegistrations) != 0 {
		err := errors.New("Invalid FCM id(s)")
		return nil, err
	}

	status, err := client.Send()

	if err != nil {
		return nil, err
	}

	return status, err
}
