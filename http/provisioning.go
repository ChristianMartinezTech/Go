package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

/*
	User Creates an Account:
		- We create a Fauna Database with that user name (Check Fauna naming conventions - dots, uppercase, numbers, special characters)
		- Whenever they create an app:
			- We create a child database
			- a Cloudlflare R2 bucket
			- a github Repo (One Repo Per APP)
	- Test each function, especially the Create Fauna Db and key creation
	We also need to create a user in a user collection in the root database of tothepoint when we create an account.
*/

// Create Github Repo function
func CreateRepoFromTemplate(appName string) {
	const baseURL = "https://api.github.com/"
	const templateOwner string = "tothepoint-app" //Before "tothepoint-app"
	const templateRepo string = "_default"
	const apiCallUrl string = baseURL + "/repos/" + templateOwner + "/" + templateRepo + "/" + "generate"
	const auth string = "token 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"

	// Struct declaration
	type templateRepoRequest struct {
		// Name is required when creating a repo.
		Name               string `json:"name,omitempty"`
		IncludeAllBranches bool   `json:"include_all_branches,omitempty"`
		Private            bool   `json:"private,omitempty"`
	}

	data := `
		{
			"name": "` + appName + `",
			"owner": "tothepoint-app",
			"private": true
		}
	`

	//body := *bytes.NewReader([]byte(data))

	body := strings.NewReader(data)
	fmt.Println(data)
	//buffer := make([]byte, 10)

	req, err := http.NewRequest("POST", "https://api.github.com/repos/"+templateOwner+"/"+templateRepo+"/generate", body)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//Reading the response
	Bodyresp, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// []Byte response to string
	respToken := string(Bodyresp)
	fmt.Println(respToken)
}

// Bucket creation
func CreateBucket(appName string) (*s3.CreateBucketOutput, error) {
	const accessKey string = "b7c03ee88daf03f9c7a25265376e5c36"
	const secretKey string = "65f8283ba4e2d31565bf132b5baf7dd24cd8d994ebc48bc5a4222e709aa1e7d0"
	const accountID string = "fbeb74f0b8ebc361d52847d7464ed188"

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.CreateBucketInput{
		Bucket: &appName,
	}

	return client.CreateBucket(context.TODO(), input)
}

func UploadWorker(appName string) {
	//const filepath = "./worker.ts"
	const filepath = "./worker.js"
	f, _ := os.Open(filepath)
	req, err := http.NewRequest("PUT", "https://api.cloudflare.com/client/v4/accounts/3a1ed33f00b15dd392fa08315b4c1e1e/workers/scripts/"+appName, f)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer l7VMKcv2Y4e2lWb46CwlrlAZ64r5CGtFvJWqc3hN") // Using Chris' Account and ID and Key
	req.Header.Set("Content-Type", "application/javascript")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//Reading the response
	Bodyresp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// []Byte response to string
	respToken := string(Bodyresp)
	fmt.Println(respToken)
}

// Create a Key-Value Namespace in Cloudflare
func CreateKVNameSpace(appName string) {
	type Payload struct {
		Title string `json:"title"`
	}

	data := Payload{appName}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.cloudflare.com/client/v4/accounts/3a1ed33f00b15dd392fa08315b4c1e1e/storage/kv/namespaces", body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer l7VMKcv2Y4e2lWb46CwlrlAZ64r5CGtFvJWqc3hN") // Using Chris' Account and ID and Key
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//Reading the response
	Bodyresp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// []Byte response to string
	respToken := string(Bodyresp)
	fmt.Println(respToken)
}

// Create FaunaDB database (from a key with admin privileges BUT will be managed from a server-level key)
// THis fucntion EXPECTS a key, if no FaunaDB server key is provided, then I'll use a key with Admin privileges
func CreateFaunaDBDatabase(app, serverKey string) {
	// Key checking in
	if serverKey == "" {
		//ENV_KEY here
		serverKey = "fnAEqxJQKMACT_jdwDsjdHhNPkRZHtVcFIjqEqQJ" //Key with admin privileges
	}

	// Fauna Client
	client := f.NewFaunaClient(
		serverKey,
	)

	// Create appName DB
	result, err := client.Query(
		f.CreateDatabase(f.Obj{"name": app}))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result) //Should this print or return?
}

// Create a key with server level access
func CreateFaunaDBServerKey(database string) (string, error) {
	// Fauna Client
	client := f.NewFaunaClient(
		"fnAEqxJQKMACT_jdwDsjdHhNPkRZHtVcFIjqEqQJ",
	)

	//Create key
	result, err := client.Query(
		f.CreateKey(
			f.Obj{
				"database": f.Database(database),
				"role":     "server",
			},
		))

	if err != nil {
		return "", err
	}
	res := result
	//fmt.Println(reflect.TypeOf(resp).Key()) // type string
	//fmt.Println(reflect.TypeOf(resp).Elem()) // type faunadb.Value
	fmt.Println(res)

	var secret string
	if err := res.At(f.ObjKey("secret")).Get(&secret); err != nil {
		return "", err
	}
	return secret, nil
}

func CreateKVPair(secret string, appName string) {

	body := strings.NewReader(secret)
	req, err := http.NewRequest("PUT", "https://api.cloudflare.com/client/v4/accounts/3a1ed33f00b15dd392fa08315b4c1e1e/storage/kv/namespaces/7b55fbe53d56444cbf13f399dec81b6c/values/FAUNADB_SECRET", body)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", "Bearer l7VMKcv2Y4e2lWb46CwlrlAZ64r5CGtFvJWqc3hN") // Using Chris' Account and ID and Key
	req.Header.Set("Content-Type", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//Reading the response
	Bodyresp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// []Byte response to string
	respToken := string(Bodyresp)
	fmt.Println(respToken)
}

func main() {
	appName := "test"
	userEmail := "testEmail@jamgo.tech"

	// User Created Account
	CreateFaunaDBDatabase(userEmail, "")

	// Create FaunaDB key
	secretUser, err := CreateFaunaDBServerKey(userEmail)
	if err != nil {
		fmt.Println(err)
	}

	// Create FaunaDB database
	CreateFaunaDBDatabase(appName, secretUser)
	// Create FaunaDB key
	secretAppDatabase, err := CreateFaunaDBServerKey(appName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(secretAppDatabase)

	// Create Repo
	//CreateRepoFromTemplate(appName)

	// Create bucket
	//CreateBucket(appName)

	// Upload Worker to R2
	//UploadWorker(appName)

	// Create a KV Namespace in Cloudflare
	//CreateKVNameSpace(appName)

	// Create Key-Value pair
	//CreateKVPair("fnAEq1MaeVACUcUEvWRsaG9TEIvzLvjGVUAj8Y", appName)
}
