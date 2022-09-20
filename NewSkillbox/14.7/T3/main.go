package main

import "fmt"

var n int

func f1(x, y int) (a, b int) {
	a = x * n
	b = y + n
	return
}

func main() {
	fmt.Print("Enter number:")
	fmt.Scan(&n)
	fmt.Print(f1(10, 10))

}
