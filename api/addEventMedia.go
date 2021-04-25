package api

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/wefolks/backend/db/query"
	"github.com/wefolks/backend/util"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
)

func EventImageHandler(client *mongo.Client) http.Handler	{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)	{
		// Max size of 100mb for the media
		maxSize := int64(1024 * 1000 * 100)
		var data  struct {
			eventId		string		`json:"event_id"`
		}

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			payload := struct {
				Error	string		`json:"error"`
			}{
				Error: "Please send the eventId",
			}
			json.NewEncoder(w).Encode(payload)
			return
		}

		event, err := query.GetEvent(data.eventId, client)

		err = r.ParseMultipartForm(maxSize)
		if err!= nil{
			// Image size too large
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			payload := struct {
				Error	string		`json:"error"`
			}{
				Error:	"Image size too large",
			}
			json.NewEncoder(w).Encode(payload)
			return
		}

		file, fileHeader, err := r.FormFile("pictures")
		if err != nil	{
			// Image size too large
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			payload := struct {
				Error	string		`json:"error"`
			}{
				Error:	"Could not upload file to server.",
			}
			json.NewEncoder(w).Encode(payload)
			return
		}
		defer file.Close()

		s3Session, err := session.NewSession(&aws.Config{
			Region: aws.String("ap-south-1"),
			Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ID"), os.Getenv("S3_SECRET"), ""),
		})

		fileName, err := util.UploadFileToS3(s3Session, file, fileHeader)

		event.PicturesUrls = append(event.PicturesUrls, fileName)
		event, err = query.UpdateEvent(event.ID, event.Name, event.Description, event.Destination, event.LocationLatitude, event.LocationLongitude, event.Datetime, event.PicturesUrls, client)
		if err != nil{
			// Image size too large
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			payload := struct {
				Error	string		`json:"error"`
			}{
				Error:	"Could not update event.",
			}
			json.NewEncoder(w).Encode(payload)
			return
		}
		payload := struct {
			_id			string		`json:"__id"`
			eventName	string		`json:"event_name"`
			pictureUrls	[]string	`json:"picture_urls"`
		}{
			_id: event.ID.String(),
			eventName: event.Name,
			pictureUrls: event.PicturesUrls,
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(payload)
		return
	})
}