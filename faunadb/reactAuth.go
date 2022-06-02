// Lambda funtion that checks the role, ID, and time stamp from a Logged in user in React Admin

package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// Lambda handler function
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	_, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return &events.APIGatewayProxyResponse{
		  StatusCode: 503,
		  Body:       "Something went wrong :(",
		}, nil
	}

******

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:      	string(jsonData),
	  }, nil
	}

func main() {
	lambda.Start(handler)	
}
