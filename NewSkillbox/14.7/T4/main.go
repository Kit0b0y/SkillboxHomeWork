package main

import "fmt"

const (
	a = 10
	b = 200
	c = 3000
)

func f1(num int) int {
	return num + a
}

func f2(num int) int {
	return num + b + c
}

func f3(num int) int {
	return num + a + b
}

func main() {
	var num int

	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	fmt.Printf("Result: %v\n", f3(f2(f1(num))))
}
