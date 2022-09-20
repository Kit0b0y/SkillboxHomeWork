package main

import "fmt"

func change(a, b int) (int, int) {
	if a > b {
		a, b = b, a
	}
	return a, b
}
func sum(a, b int) int {
	if a&1 != 0 {
		a++
	}
	if b&1 != 0 {
		b--
	}

	var n = (b-a)/2 + 1
	return (a + b) * n / 2
}
func main() {
	var a, b int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Printf("Сумма четных чисел в диапазоне равна %v", sum(a, b))
}
