package util

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

const	(
	MESSAGE	=	"$OTP$ is your OTP to verify your phone number on the folks application."
	url	=	"https://www.fast2sms.com/dev/bulkV2"
)
func generateOTP() string	{
	num := 1000+rand.Intn(9999-1000)
	return string(num)
}

// Function to send OTP, OTP is NA in case of error, else OTP is returned
func SendOTP(n int)	(string, error){
	otp := generateOTP()
	message := strings.Replace(MESSAGE, "$OTP$", otp, 1)
	payload,_ := json.Marshal(map[string]string{
		"authorization": "Vl5Unv6pTqWsZI4zR1rDxy0b9PuOYGjHJeE3i2FfABcQaXSgkwKDApEWtI2Z1RaVGNbrl59i87yJPSFc",
		"message":	message,
		"language": "english",
		"route": "q",
		"numbers": string(n),
	})
	responseBody :=	bytes.NewBuffer(payload)
	_, err := http.Post(url, "application/json", responseBody)
	if err!=nil	{
		log.Fatalf("Unable to send message")
		return "#NA", err
	}
	return otp, nil
}