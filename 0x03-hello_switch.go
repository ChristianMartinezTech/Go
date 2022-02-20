// Guessing number program
// Program implementing switch and rand

package main

import (
	"fmt"
	"math/rand"
)

func switcher(guess int) {
	// Getting the rand number range
	n := (rand.Intn(2)) + 1

	// Checking it's a valid guess
	switch guess > 3 || guess <= 0 {
	case true:
		fmt.Println("That's an invalid number.")
		fmt.Println("Pick one, two, or three.")
		break

	case false:
		if guess == n {
			fmt.Println("Bingo!!! the number is", n)
		} else {
			fmt.Println("Try again. the number is", n)
		}
	}
}

func main() {
	// Getting the guess
	fmt.Println("This is a small number game.")
	fmt.Println("Is it 1, 2 or 3? Guess Bellow:")
	var guess int
	fmt.Scanln(&guess)

	// Function call
	switcher(guess)
}
