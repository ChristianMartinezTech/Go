// Server that can make GET, POST, DELETE, and PATCH http request to the baserow API
// Test table https://baserow.io/database/26818/table/59764

// GET: "curl -H "Authorization: Token zckM7icpJJjnCDgDJjrCFx57diqFYdbZ" https://api.baserow.io/api/database/fields/table/59764/"

package main

import (
	"fmt"
	"log"
	"net/http"
)

var Token = "Authorization: Token zckM7icpJJjnCDgDJjrCFx57diqFYdbZ"
var TableId = 59764
var Url = "https://api.baserow.io/api/database/fields/table/59764/"

func main() {
	// 0. Check on the Baserow authentication

	// 1. Instantiate the Go Client so we can add the "Authorization: Token YOUR_API_KEY" header
	client := &http.Client{}

	// 2. Construir los requests
	req, err := client.Get("https://api.baserow.io/api/database/fields/table/59764/")
	req.Header.Set(Token, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req)
}
