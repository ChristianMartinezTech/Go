package main

import (
	"fmt"
	"time"
)

func Time() {
	Now := time.Now()
	fmt.Println(Now)
}

func main() {
	Time()
}
