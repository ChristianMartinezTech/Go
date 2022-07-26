package main

import (
	"context"
	"fmt"

	//"os"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"

	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func createRoleWithPerm() string {
	client := f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")
	result, err := client.Query(
		f.CreateRole(
			f.Obj{
				"name": "newrolejueves",
				"privileges": f.Obj{
					"resource": f.Collection("managers"),
					"actions":  f.Obj{"read": true}}}))
	if err != nil {
		log.Fatal(err)
	}
	RoleWithPermCreated := fmt.Sprintf("%v", result)
	//fmt.Println(result)
	return RoleWithPermCreated
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

	result := createRoleWithPerm()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       result,
	}, nil
}

func main() {
	//createRoleWithPerm()
	lambda.Start(handler)
}
