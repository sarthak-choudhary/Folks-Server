package util

import (
	"bytes"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"gopkg.in/mgo.v2/bson"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func uploadFileToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	tempFileName := bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String("we-folks"),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tempFileName, err
}

// Handler for Image uploads to aws
func Handler(w http.ResponseWriter, r *http.Request) {
	maxSize := int64(1024 * 1000 * 100) // allow 100MB of file size

	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		//Image size too large
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		payload := struct {
			Error string `json:"error"`
		}{Error: "Image size more than 100MB"}

		json.NewEncoder(w).Encode(payload)
		return
	}

	file, fileHeader, err := r.FormFile("pictures")
	if err != nil {
		//Could not get uploaded file
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		payload := struct {
			Error string `json:"error"`
		}{Error: "Could not get uploaded file."}

		json.NewEncoder(w).Encode(payload)
		return
	}
	defer file.Close()

	s, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(
			"AKIA4JUME47GK5B7FMKX",
			"e19XAPTiy2ktiFfau3NKWVdRm716USG9Ge5z6oCF",
			"",
		),
	})

	if err != nil {
		//Could not create session with S3
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fileName, err := uploadFileToS3(s, file, fileHeader)
	if err != nil {
		//Could not upload file to S3
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload := struct {
		FileName string `json:"filename"`
	}{
		FileName: fileName,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payload)
}
