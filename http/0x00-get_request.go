// Simple hhtp request program -> go run 0x00-get_request.go {url}
// it will receive 1 parameter as the url to make the request to

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Get request funct
func api(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func main() {
	url := os.Args[1]
	api(url)
}
