package main

import (
//	"os"
	"reflect"
	f "github.com/fauna/faunadb-go/v4/faunadb"
	"encoding/json"
 	"context"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda" 
) 

//const faunaSecret = os.Getenv("FAUNA_SECRET") 

const faunaSecret = "fnAEn2lG7lACQRG7CMYhotZjga9243wim8F1vz1o"

type Customer struct {
	ID		string	`json:"id"`
	FirstName string `fauna:"firstName" json:"first_name"`
	LastName string `fauna:"lastName" json:"last_name"`
	Address interface{} `fauna:"address" json:"address"`
	Telephone string `fauna:"telephone" json:"phone"`
}

func GetIDS(res f.Value) []string {
	var refs []f.RefV
	var ids []string
	if err := res.At(f.ObjKey("ref")).Get(&refs); err != nil {
		panic(err)
	}
	for _,ref := range refs {
		ids = append(ids, ref.ID)
	}
	return ids
}

func GetMany(e interface{}, res f.Value) string {
	T := reflect.TypeOf(e)
	if T.Kind()!= reflect.Slice {
		return ""
	}
	x := reflect.New(T).Interface()
	if err := res.At(f.ObjKey("data")).Get(&x); err != nil {
		panic(err)
	}
	ids := GetIDs(res)
	reflect.ValueOf(x).Elem().FieldByName("ID").SetString(id)
	json,_:= json.Marshal(x)
	return string(json)
}


func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	_, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return &events.APIGatewayProxyResponse{
		  StatusCode: 503,
		  Body:       "Something went wrong :(",
		}, nil
	} 
	
	client := f.NewFaunaClient(faunaSecret)

	ids := request.MultiValueQueryStringParameters["ids"]

	var queries []f.Expr 

	for _,id := range ids {
		queries = append(queries, `Get(Ref(Collection("managers"), id)))`)
	}
		
	res, err := client.BatchQuery(queries)
	if err != nil {
		panic(err)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       GetMany([]Customer{},res),
	}, nil
}

func main() {
	lambda.Start(handler)
} 



