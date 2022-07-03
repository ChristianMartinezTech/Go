package main

import (
	"fmt"
)

type figure2d interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	height float64
	width  float64
}

func (s square) area() float64 {
	return s.base * s.base
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func calculate(f figure2d) {
	fmt.Println("Area:", f.area())
}

func main() {
	mySquare := square{base: 4}
	myRectangle := rectangle{height: 3, width: 5}

	calculate(mySquare)
	calculate(myRectangle)
}
