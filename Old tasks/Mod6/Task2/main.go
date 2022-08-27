package main

import (
	"fmt"
)

func main() {

	userInput1 := 0
	userInput2 := 0
	fmt.Println("Enter number1:")
	fmt.Scan(&userInput1)
	fmt.Println("Enter number2:")
	fmt.Scan(&userInput2)
	sum := userInput1 + userInput2

	i := 0

	for {
		i++
		if i == sum {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("Сумма:", sum)
}
