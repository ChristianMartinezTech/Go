// Program that iterate over two numbers given by the use
// Used if/else for cycles, and function calls

package main

import "fmt"

func iterationFunction(numberOne, numberTwo int) {

	// Determing which number if bigger
	if numberOne >= numberTwo {

		// Printing with for Cycle
		fmt.Println("First Number is bigger:")
		for i := numberTwo; i <= numberOne; i++ {
			fmt.Println(i)
		}
	} else {

		// Printing with for Cycle
		fmt.Println("Second Number is bigger:")
		for i := numberOne; i <= numberTwo; i++ {
			fmt.Println(i)
		}
	}
}

func main() {
	// Getting the two numbers from the user
	fmt.Println("Please input two numbers:")
	var numberOne int
	var numberTwo int
	fmt.Scanln(&numberOne, &numberTwo)

	// Function call
	iterationFunction(numberOne, numberTwo)
}
