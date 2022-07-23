package main

import (
	"fmt"
)

func main() {
	fmt.Println("Проверка положительного числа")
	var num1 int
	var num2 int
	var num3 int
	fmt.Println("Введите первое число:")
	fmt.Scan(&num1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&num2)
	fmt.Println("Введите третье число:")
	fmt.Scan(&num3)

	if num1 > 0 || num2 > 0 || num3 > 0 {
		fmt.Println("Одно из введенных чисел положительное.")
	} else {
		fmt.Println("Ни одно из введенных чисел  не является положительным.")
	}
}
