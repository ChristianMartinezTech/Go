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

func getDocument() string {
	client := f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")

	res, err := client.Query(f.Get(f.Ref(f.Collection("managers"), "101")))
	if err != nil {
		log.Fatal(err)
	}
	var customer Customer

	if err := res.At(f.ObjKey("data")).Get(&customer); err != nil {
		log.Fatal(err)
	}
	getDocument := fmt.Sprintf("%v", customer)
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

	result := getDocument()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       result,
	}, nil
}

func main() {
	//getDocument()
	lambda.Start(handler)
}
