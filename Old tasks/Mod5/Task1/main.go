package main

import (
	"fmt"
)

func main() {
	fmt.Println("Определение координатной плоскости.")

	var x int
	var y int
	fmt.Println("Введите координту Х:")
	fmt.Scan(&x)
	fmt.Println("Введите координту Y:")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Println("Точка принадлежит 1 координтной четверти")
	}
	if x < 0 && y > 0 {
		fmt.Println("Точка принадлежит 2 координтной четверти")
	}
	if x < 0 && y < 0 {
		fmt.Println("Точка принадлежит 3 координтной четверти")
	}
	if x > 0 && y < 0 {
		fmt.Println("Точка принадлежит 4 координтной четверти")
	}
}
