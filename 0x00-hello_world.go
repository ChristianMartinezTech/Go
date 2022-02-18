package main

// fmt printintg library
import (
	"fmt"
	"os"
)

func main() {
	// var declaration
	message := "Hello,"
	// arg0 := os.Args[0] will be the program path
	arg1 := os.Args[1] + "."
	const message2 = "Have a good one."

	fmt.Println(message, arg1, message2)
	fmt.Println("PD: Hello World!")
}
