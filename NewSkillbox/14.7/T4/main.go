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
	return num + b
}

func f3(num int) int {
	return num + c
}

func main() {
	var num int

	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	fmt.Printf("Result 1 (+10): %v\n", f1(num))
	fmt.Printf("Result 2 (+200): %v\n", f2(num))
	fmt.Printf("Result 3 (+3000): %v\n", f3(num))

	fmt.Printf("Result (+ 10)+200)+3000): %v\n", f3(f2(f1(num))))
}
