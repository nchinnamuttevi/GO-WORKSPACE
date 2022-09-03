package main

import (
	"fmt"
)

func main() {
	var x float64
	var result float64
	fmt.Println("What value do you need the square of")
	fmt.Scan(&x)
	result = x * x
	fmt.Printf("the square of the value: %d you entered is %d", x, result)
}
