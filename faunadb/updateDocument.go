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

func updateDocument() string {
	client := f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")
	res, err := client.Query(
		f.Update(f.Ref(f.Collection("managers"), "101"), f.Obj{"data": f.Obj{"name": "Mountain's Thunder"}}))
	if err != nil {
		log.Fatal(err)
	}

	var customer Customer
	if err := res.At(f.ObjKey("data")).Get(&customer); err != nil {
		log.Fatal(err)
	}
	documentUpdated := fmt.Sprintf("%v", res)
	//fmt.Println(customer)
	return documentUpdated
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
	result := updateDocument()
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       result,
	}, nil
}

func main() {
	//updateDocument()
	lambda.Start(handler)
}
