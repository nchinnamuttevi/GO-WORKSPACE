package main

import (
	"fmt"
)

func main() {
	var TableNumber int
	var TableLength int
	var CurrentVal int
	i := 0
	fmt.Println("What multiplication table do you want")
	fmt.Scan(&TableNumber)
	fmt.Println("Upto how many do you want the numbers")
	fmt.Scan(&TableLength)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	for {
		if i <= TableLength {
			CurrentVal = TableNumber * i
			fmt.Printf("%d X %d = %d \n", TableNumber, i, CurrentVal)

		} else {
			break
		}
		i = i + 1
	}
}
