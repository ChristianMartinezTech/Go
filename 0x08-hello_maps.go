// file to practice creating and handling maps

package main

import (
	"fmt"
)

func price() {
	menu := map[string]int{
		"apple":        1,
		"pinapple":     5,
		"coconut":      3,
		"strawberries": 2,
		"grapes":       3,
	}
	fmt.Println(menu)

	// Taking the customer's order
	openingMsg := `Hey there! Here's today's menu:
	apple
	pinneaple
	coconut
	strawberries
	grapes`
	fmt.Println(openingMsg)
	fmt.Println("What would you like to order?")

	var order string
	fmt.Scanln(&order)
	//fmt.Println("Your order is: " + order)

	//Check if order is in the menu
	if menu[order] != nil {

	}
}

func main() {
	price()
}
