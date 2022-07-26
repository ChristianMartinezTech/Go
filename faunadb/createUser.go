package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

var client *f.FaunaClient

// Fauna client
func getToken() string {
	client = f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")
	// Example on logging in
	result, err := client.Query(
		f.Select("secret",
			f.Login(
				f.Ref(f.Collection("users"), "333401868560499279"),
				f.Obj{"password": "secretpassword"},
			)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	str := fmt.Sprintf("%v", result)
	return (str)
}

// Function to look for users by ref
func indexSearch(string) {

}

// Function that checks pemission
func permisions() {
	// get permissions of an specific role
	result, err := client.Query(
		f.SelectAll("actions", f.Select("privileges",
			f.Get(f.Role("new-role4")))))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(result)
	}
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

	getToken()
	permisions()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Not finished yet", // Finish this func
	}, nil
}

func main() {
	//lambda.Start(handler)
	getToken()
	//permisions()
}
