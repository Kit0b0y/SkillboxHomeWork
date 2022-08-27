package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ёлочка")
	fmt.Println("Введите высоту ёлки:")
	var height int
	fmt.Scan(&height)

	for i := 1; i <= height; i++ {
		for j := 0; j < height-i; j++ {
			print(" ")
		}
		for j := 0; j < i*2-1; j++ {
			print("*")
		}
		fmt.Println()
	}
}
