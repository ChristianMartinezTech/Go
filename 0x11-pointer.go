// Simple program to test ponter and deferencing

package main

import "fmt"

func main() {
	var var1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", var1, &var1)

	var var1Pointer *int
	var1Pointer = &var1

	fmt.Printf("Pointer value: %d, it's location in memory: %p\n", var1Pointer, &var1Pointer)

	*var1Pointer = 10
	fmt.Printf("Deferenced integer value: %d, it's location in memory: %p\n", var1, &var1)
}
