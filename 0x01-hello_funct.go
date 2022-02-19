// Program that determines if the user imput is even or odd
// Used if else statements, asking for user input and function calls

package main

import (
	"fmt"
)

// Function to determine if number is even or odd
func evenOrOdd(numbr int) {
	if numbr%2 == 0 {
		fmt.Println(numbr, "is even!")
	} else {
		fmt.Println(numbr, "is odd!")
	}
}

func main() {
	// Taking in user input
	fmt.Println("Please input a number:")
	var numbr int
	fmt.Scanln(&numbr)

	//Function call
	evenOrOdd(numbr)
}
