package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"

	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func createCollection() string {
	client := f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")

	res, err := client.Query(f.CreateCollection(f.Obj{"name": "boons"}))
	if err != nil {
		log.Fatal(err)
	}
	collectionCreated := fmt.Sprintf("%v", res)
	//fmt.Println(collectionCreated)
	return (collectionCreated)
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

	collectionCreated := createCollection()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       collectionCreated,
	}, nil
}

func main() {
	//createCollection()
	lambda.Start(handler)
}
