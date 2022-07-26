package main

import (
	"context"
	"fmt"
	"log"

	f "github.com/fauna/faunadb-go/v4/faunadb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func getPermissionsSpecificRole() string {
	client := f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")

	result, err := client.Query(
		f.SelectAll("actions", f.Select("privileges",
			f.Get(f.Role("new-role4")))))
	if err != nil {
		log.Fatal(err)
	}

	getDocument := fmt.Sprintf("%v", result)
	//fmt.Println(getDocument)
	return getDocument
}

// lambda function
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	_, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return &events.APIGatewayProxyResponse{
			StatusCode: 503,
			Body:       "Something went wrong :(",
		}, nil
	}

	getDocument := getPermissionsSpecificRole()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       getDocument,
	}, nil
}

func main() {
	//getPermissionsSpecificRole()
	lambda.Start(handler)
}
