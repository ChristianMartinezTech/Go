// Method that checks the validy of the JWS Authentication and GetLists

package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-shiori/go-readability"
)

// Customer
type Customer struct {
	FirstName string `fauna:"firstName"`
	LastName  string `fauna:"lastName"`
	//Address   fields `fauna:"address"`
	Telephone string `fauna:"telephone"`
}

// Event
type Event struct {
	Url string `json:"url"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var event Event
	json.Unmarshal([]byte(request.Body), &event)
	article, err := readability.FromURL(event.Url, 300*time.Second)
	if err != nil {
		panic(err)
	}

	return events.APIGatewayProxyResponse{Body: article.Content, StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
