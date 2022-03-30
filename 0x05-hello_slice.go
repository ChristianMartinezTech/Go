// Program to use slices
// In go arrays have a fixed lenght

package main

import "fmt"

func main() {
	slice := []string{"I", "am", "your", "father"}

	// for _, valor := range slice {
	// 	fmt.Printf(valor + " ")
	// }
	// fmt.Printf("\n")

	for i, v := range slice {
		fmt.Println(i, v)
	}
}
