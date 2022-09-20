package main

import (
	"fmt"
)

func firstFunc(a *int) bool {
	if *a%2 == 0 {
		return true
	}
	return false
}

func main() {
	var n int
	print("Введите число:")
	fmt.Scan(&n)

	if firstFunc(&n) {
		fmt.Printf("Число %v - четное\n", n)
	} else {
		fmt.Printf("Число %v - нечетное\n", n)
	}
}
