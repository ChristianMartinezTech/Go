package main

import (
	"context"
	"fmt"
	"log"

	// "reflect"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"

	f "github.com/fauna/faunadb-go/v4/faunadb"
)

// Initialize and apply formula we already have
/*func ApplyFormulas(e interface{}) {
	T := reflect.TypeOf(e)
	if T.Kind()!= reflect.Struct {
		return
	}
	for i:=0; i<T.NumField(); i++ {
		ft := T.Field(i)
		formulaName := ft.Tag(formula)
		if formulaName="" {
			return
		}
		inputsNames := ft.Tag(inputs)
		var inputs []interface{}
		for _,in := range inputsNames {
			inputs = append(inputs, reflect.ValueOf(e).FieldByName(in).Elem())
		}
		output := formulas[formulaName](inputs...)
		fv := reflect.ValueOf(e).Field(i)
		if !(fv.CanSet) {
			return
		}
		fv.Set(reflect.ValueOf(output))
}
}*/

func createDoc() {
	client := f.NewFaunaClient("fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o")
	res, err := client.Query(
		f.Create(
			f.Collection("users"),
			f.Obj{"data": f.Obj{"name": "paula"}},
		))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
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

	createDoc()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Document created",
	}, nil
}

func main() {
	lambda.Start(handler)
}
