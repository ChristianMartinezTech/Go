package main

import (
	//"fmt"
	"log"

	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"

	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func deleteDocument() string {
	client := f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")

	res, err := client.Query(f.Delete(f.Ref(f.Collection("users"), "333840942921417281")))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	deleteDocument := fmt.Sprintf("%v", res)
	return deleteDocument
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

	result := deleteDocument()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       result,
	}, nil
}

func main() {
	//deleteDocument()
	lambda.Start(handler)
}
