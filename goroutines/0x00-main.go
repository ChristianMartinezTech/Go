// First experiment with Go routines and concurrency :)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func showGoRoutine(id int) {
	delay := rand.Intn(500)
	fmt.Printf("Go Routine #%d with %d\n", id, delay)

	time.Sleep(time.Millisecond * time.Duration(delay))
}

func main() {
	// Lopping to execute the Go routine
	for i := 0; i <= 20; i++ {
		showGoRoutine(i)
	}
}
