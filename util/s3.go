package util

import (
	"bytes"
	"fmt"
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

	tempFileName := "pictures/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

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
		log.Println(err)
		fmt.Fprintf(w, "Image too large. Max Size: %v", maxSize)
		return
	}

	file, fileHeader, err := r.FormFile("pictures")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "could not get uploaded file")
		return
	}
	defer file.Close()

	s, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(
			"AKIA4JUME47GBZXQ76VW",
			"kH4BdpHoOttHJk88Nf9I5u4yuZXJUw06ATnaGZ6A",
			"",
		),
	})
	if err != nil {
		fmt.Fprintf(w, "Could not create session")
	}

	fileName, err := uploadFileToS3(s, file, fileHeader)
	if err != nil {
		fmt.Fprintf(w, "Could not upload file to s3")
	}

	fmt.Fprintf(w, "Image uploaded successfully: %v", fileName)
}
