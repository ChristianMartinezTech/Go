package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"

	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func getMany() {
	client := f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")

	// Code here
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

	getMany()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Document created",
	}, nil
}

func main() {
	lambda.Start(handler)
}
