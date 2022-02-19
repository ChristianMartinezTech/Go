// Program to use args and multiple ways of printing
// Usage: go run 0x00-hello_var.go [Name]

package main

// import args and fmt
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
	const old_programming_lang = "C"

	// To initialize empty var
	// var a int
	// var b string
	// var c float64
	// var d bool

	fmt.Println(message, arg1, message2)
	fmt.Println("Let's print a string, sum a var and a value:", number-5)
	fmt.Printf("Now let's use Printf like in %s\n", old_programming_lang)
}
