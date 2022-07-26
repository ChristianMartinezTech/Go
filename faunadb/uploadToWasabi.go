package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// How to set the credentials
const accessKey string = "UQRZBUJQ7SG7UX7MVFPB"
const secretKey string = "sb1jRQzEhJC8giHWkWO5OY20D64blZzYN30qmiEo"
const wasabiBucket string = "examplejamgobucket"

var filePath string = "/home/chriswin/JamgoTech/wasabi/21st-century-c-o-reilly-ben-klemens.pdf"

func uploadToWasabi() {
	// create a configuration for profile
	s3Config := aws.Config{
		Endpoint:         aws.String("s3.eu-central-1.wasabisys.com"),
		Region:           aws.String("eu-central-1"),
		S3ForcePathStyle: aws.Bool(true),
	}

	// create a new session using the config above and profile
	goSession, err := session.NewSessionWithOptions(session.Options{
		Config:  s3Config,
		Profile: "wasabi",
	})
	if err != nil {
		log.Fatal(err)
	}

	// create a s3 client session
	s3Client := s3.New(goSession)
	fmt.Println(s3Client)

	//set the file path to upload
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//create put object input
	putObjectInput := &s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(wasabiBucket),
		Key:    aws.String("text.txt"), //Still having to fix this
	}

	// upload file
	_, err = s3Client.PutObject(putObjectInput)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	uploadToWasabi()
}
