package main

// fmt printintg library
import (
	"fmt"
	"os"
)

func main() {
	// var declaration
	message := "Hello,"
	arg1 := os.Args[1] + "."
	const message2 = "Have a good one."
	number := 10

	// To initialize empty var
	// var a int
	// var b string
	// var c float64
	// var d bool

	fmt.Println(message, arg1, message2)
	fmt.Println("PD: Hello World!")
	fmt.Println(number - 5)
}
