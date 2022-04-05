// Program that makes a post request to the airtable api
// https://docs.aws.amazon.com/es_es/lambda/latest/dg/golang-handler.html

// To add favorites in the browser -> cheks if the user is loged in -> 
// triggers lambda function (listen to the AIR call from the front end) -> 
// makes post request to the airtable API

package main

import (
	//"fmt"
	"http"
	//"context"
	"github.com/aws/aws-lambda-go/lambda"
)

func postRequest {
	// Post request to the airtable api
	resp, err := http.PostForm("https://api.airtable.com/v0/appnGSgUk1lPoc5oX/test",
	url.Values{"key": {"Value"}, "id": {"123"}})

}

func main {
	// Start postRequest func
	lambda.Start(postRequest)
}
