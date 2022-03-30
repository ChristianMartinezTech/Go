// Program that adds all the input in an array
// Used arrays, note that slices is preffered over arrays

// Unlike in C/C++ (where arrays act like pointers) and Java
// (where arrays are object references), arrays in Go are values.
// This has a couple of important implications: (1) assigning one array to another copies all of the elements
// (2) if you pass an array to a function, it will receive a copy of the array (not a pointer or reference to it).

// As you might imagine, this can be very expensive, especially when you are working with arrays that have a large number of elements.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Taking in the size of the array
	fmt.Println("Hi there! This program saves an array of numbers")
	fmt.Println("What's the size of the array?")
	var size int
	fmt.Scanln(&size)

	// Taking in the array values
	fmt.Println("Thank you. Please input your sequence of numbers:")
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", (i + 1))
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Your array is: ", arr)
	fmt.Println("Its datatype is: ", reflect.TypeOf(arr))
}
