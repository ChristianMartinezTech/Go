// Code to practice short variable declaration

package main

import "fmt"

func main() {
	var i, j int = 1, -1
	var q int
	l := "hola"
	m, n, o := true, false, "go"

	fmt.Printf("Value:%v Var type: %T\n", i, i)
	fmt.Printf("Value:%v Var type: %T\n", j, j)
	fmt.Printf("Value:%v Var type: %T\n", q, q)
	fmt.Printf("Value:%v Var type: %T\n", l, l)
	fmt.Printf("Value:%v Var type: %T\n", m, m)
	fmt.Printf("Value:%v Var type: %T\n", n, n)
	fmt.Printf("Value:%v Var type: %T\n", o, o)
}
