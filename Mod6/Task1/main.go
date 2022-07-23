package main

import (
	"fmt"
)

func main() {
	i := 0
	userInput := 0
	fmt.Println("Enter number:")
	fmt.Scan(&userInput)

	for {
		i++
		if i == userInput {
			break
		}
		fmt.Println(i)
	}
}
