// Calculator. Use: go run calculator.go [operation]

package main

import (
	"os"
)

// this func separates and cast the input from str to int
func operation(UsrInput string) {
	for i := 0; i < len(UsrInput); i++ {

	}

}

//func operation()

func main() {
	// Taking the user input and send it to operation
	UsrInput := os.Args[1]
	/*fmt.Println(UsrInput)
	fmt.Printf("%T\n", UsrInput)*/
	operation(UsrInput)
}
