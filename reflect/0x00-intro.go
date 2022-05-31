// Reflection introduction

package main

import (
	"fmt"
	//"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	var x float64 = 3.14
	var u User = User{"bob", 10}

	fmt.Println(x)
	fmt.Println(u)

}

/* Questions to respond:
- What's "type"


*/