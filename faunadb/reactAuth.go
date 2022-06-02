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

	/*The employee is logged in React Admin
	Check the Json token - Should contain the role, ID, and time stamp

	1. Lambda that checks the token
	- if its expired, "Error - Token has expired"
	- Use the role scope "Error - Permision denied"

	2. Get the role and ID of the employee that is inside the Json token
	- Use them to query to the fauna db
	- Show the information that is relevant to the employee ID

	- React Admin uns the Auth provider
	- It's sent to the auth server in Fauna
	- If the credentials are correct the auth provider returns the token (check what info has)
	- Create 1 login enpoint to each role*/

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonData),
	}, nil
}

func main() {
	lambda.Start(handler)
}
