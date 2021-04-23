package fcm

import (
	fcm "github.com/douglasmakey/go-fcm"
	"os"
)

//const (
//	serverKey = "AAAA57PWsLA:APA91bFoZewaBJImbmxONlL_sNW5TS2138Om_IQIsXt53TiHnDG1HOQHo8hUDLlqFC30dCUCnywFkWcsiXu7qXP2Dyt8jDBty82ZNZSXTIvUFErRjrCWNEmxtwwt_022gg5wYrn9m_Js"
//)

//GetFcmClient - gives the connection to Firebase
func GetFcmClient() *fcm.Client {
	return fcm.NewClient(os.Getenv("FCM_SERVER_KEY"))
}
