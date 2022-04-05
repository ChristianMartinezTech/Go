// Simple hhtp server in GO

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Creating the Json struct
type Order struct {
	Product string "json:'product'"
	Amount  int    "json:'amount'"
	Price   int    "json:'price'"
}

// Serving Json Order
func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	order := Order{
		Product: "Apple Macbook air M1",
		Amount:  1,
		Price:   1000,
	}
	json.NewEncoder(w).Encode(order)
}

// Hello testing func
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello there!"))
}

// http serving
func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/order", GetOrder)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
