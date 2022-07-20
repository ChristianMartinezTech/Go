package main

import (
	"fmt"
)

func main() {
	a := 10
	b := &a
	c := *b

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
