package api

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/wefolks/backend/db/models"
	"github.com/wefolks/backend/db/query"
	"github.com/wefolks/backend/util"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
)

func ChangeProfilePicture(client *mongo.Client) http.Handler	{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)	{
		// Max size of 100mb for the media
		maxSize := int64(1024 * 1000 * 100)
		user := r.Context().Value("user").(*models.User)

		err := r.ParseMultipartForm(maxSize)
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
		if user.ProfilePicture == ""{
			user.ProfilePicture = fileName
		}	else {
			user.PicturesUrls = append(user.PicturesUrls, user.ProfilePicture)
			user.ProfilePicture = fileName
		}
		_, err = query.UpdateUser(user, client)
		if err != nil {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			payload := struct {
				Error	string		`json:"error"`
			}{
				Error:	"Could not update user",
			}
			json.NewEncoder(w).Encode(payload)
			return
		}

		payload := struct {
			_id					string		`json:"_id"`
			ProfilePicture		string		`json:"profilePicture"`
		}{
			_id: user.ID.String(),
			ProfilePicture: user.ProfilePicture,
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(payload)
		return
	})
}