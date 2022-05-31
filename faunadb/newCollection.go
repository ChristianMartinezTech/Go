// Program that creates a new collection called "myCollection"

package main

import (
	"fmt"
	"os"

	f "github.com/fauna/faunadb-go/v4/faunadb"
)

// Acquire the secret and optional endpoint from environment variables
var (
	secret   = "fnAEn7ciCsACUfzmdzyHPeDAmr9FJ4_ocnrKfKg_"
	endpoint = "https://db.fauna.com"
)

func main() {
	// Instantiate a client
	client := f.NewFaunaClient(secret, f.Endpoint(endpoint))

	// Create a collection called 'myCollection'
	result, err := client.Query(
		f.CreateCollection(f.Obj{"name": "myCollection"}))

	// Show the output
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(result)
	}
}
