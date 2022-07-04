// Calculator. Use: go run calculator.go [operation]

package main

import (
	"fmt"
	"log"
	"strconv"
)

// Function that converts a string into int
func operation(UsrInput string) {

	for i := 0; i < len(UsrInput); i++ {
		op, err := strconv.Atoi(UsrInput)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(op)
	}
}

func experiment() {
	a := 2
	b := "+"
	c := 5

	fmt.Printf("%T/n", a)
	fmt.Printf("%T/n", b)
	fmt.Printf("%T/n", c)
	fmt.Println(a, b, c)

}

func main() {
	// Taking the user input and send it to operation
	/*fmt.Println("Hello there! Input your operation:")
	var UsrInput string
	fmt.Scanf("%s", &UsrInput)
	operation(UsrInput)*/

	experiment()

}
