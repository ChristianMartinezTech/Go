// file to practice creating and handling maps

package main

import "fmt"

func price() {
	// Taking the customer's order
	openingMsg := "Hey there! Here's today's menu:
	apple
	pinneaple
	coconut
	strawberries
	grapes"

	fmt.Println()
	fmt-Println("What would oyu like to order?")
	var order string
	fmt.Scanf(&order)

	// Detecting if the client response is in the menu
	var dict map[string]int = map[string]int {
		"apple": 1,
		"pinapple": 5,
		"coconut": 3,
		"strawberries": 2,
		"grapes": 3
	}

	

func main() {


	}
}