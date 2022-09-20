package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a, b int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Printf("Значения до вызова функции: %v и %v\n", a, b)
	swap(&a, &b)
	fmt.Printf("Значения после вызова функции: %v и %v\n", a, b)
}
