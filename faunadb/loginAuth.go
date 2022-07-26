package main

import (
	"context"
	"fmt"
	"os"

	f "github.com/fauna/faunadb-go/v4/faunadb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type fields struct {
	Street  string `fauna:"street"`
	City    string `fauna:"city"`
	State   string `fauna:"state"`
	ZipCode string `fauna:"zipCode"`
}

type Customer struct {
	FirstName string `fauna:"firstName"`
	LastName  string `fauna:"lastName"`
	Address   fields `fauna:"address"`
	Telephone string `fauna:"telephone"`
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

	// Code
	client := f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")
	result, err := client.Query(
		f.Login(
			f.Ref(f.Collection("managers"), "101"),
			f.Obj{"password": "abracadabra"},
		))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(result)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(result), //Check this one out
	}, nil
}

func main() {
	lambda.Start(handler)
}
