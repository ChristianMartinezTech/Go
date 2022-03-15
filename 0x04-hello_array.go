// Program that adds all the input in an array
// USed arrays and slices

package main

import "fmt"

func arr(user_array []int) {
	fmt.Println("Great! This is your number sequence:")

	for i, s := range user_array {
		fmt.Println(i, s)
	}

}

func main() {
	// Function opening messages
	fmt.Println("Hi there! This program saves a sequence of numbers")
	fmt.Println("You can review the position of each umber in a sequence")
	fmt.Println("Please input your sequence of numbers:")
	var user_array []int
	fmt.Scanln(&user_array)

	// Function call
	arr(user_array)
}
