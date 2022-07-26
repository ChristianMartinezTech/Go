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

type Customer struct {
	Name string `fauna:"name"`
}

func changeNametoCollection() string {
	// Code
	client := f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")
	res, err := client.Query(f.Update(f.Collection("products2"), f.Obj{"name": "boons"}))
	if err != nil {
		log.Fatal(err)
	}

	var customer Customer
	if err := res.Get(&customer); err != nil {
		log.Fatal(err)
	}

	nameOfCollectionChanged := fmt.Sprintf("%v", res)
	//fmt.Println(nameOfCollectionChanged)
	return nameOfCollectionChanged
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
	result := changeNametoCollection()
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       result,
	}, nil
}
func main() {
	//changeNametoCollection()
	lambda.Start(handler)
}
