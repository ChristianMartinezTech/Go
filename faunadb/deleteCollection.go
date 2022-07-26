package main

import (
	//"fmt"

	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func deleteCollection() string {
	client := f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")

	res, err := client.Query(f.Delete(f.Collection("boons2")))
	if err != nil {
		log.Fatal(err)
	}
	collectionDeleted := fmt.Sprintf("%v", res)
	//fmt.Println(collectionDeleted)
	return (collectionDeleted)
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

	result := deleteCollection()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       result,
	}, nil
}

func main() {
	//deleteCollection()
	lambda.Start(handler)
}
