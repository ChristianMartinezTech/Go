package main

import (
	"fmt"
	"log"
)

func suma() {
	a := 5
	b := 3
	c := "Holaaa"

	suma1, err := a + b
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(suma1)

	suma2, err := b + c
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(suma2)
}

func main() {
	suma()
}
