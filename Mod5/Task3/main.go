package main

import (
	"fmt"
)

func main() {
	fmt.Println("Проверка совпадающих чисел")
	var num1, num2, num3 int

	fmt.Println("Введите первое число:")
	fmt.Scan(&num1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&num2)
	fmt.Println("Введите третье число:")
	fmt.Scan(&num3)

	if num1 == num2 || num2 == num3 || num3 == num1 {
		fmt.Println("Одно или несколько введенных чисел совпадают.")
	} else {
		fmt.Println("Ни одно из введенных чисел  не совпадают друг с другом.")
	}
}
