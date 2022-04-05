// function that detects if a word is palindrome

package main

import "fmt"

func is_palindrome(text string) {
	var reversed_text string

	// Creating a variable with the string reversed
	for i := len(text) - 1; i >= 0; i-- {
		reversed_text += string(text[i])
	}

	// Checking if its palindrome
	if reversed_text == text {
		fmt.Println("The word " + text + "is palindrome!")
	} else {
		fmt.Println("The word " + text + "is NOT palindrome")
	}
}

func main() {
	fmt.Println("Hi there! Let us check if your word is palindrome:")
	var text string
	fmt.Scanf("%s", &text)

	// function call
	is_palindrome(text)
}
