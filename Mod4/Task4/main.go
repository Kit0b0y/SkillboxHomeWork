package main

import (
	"fmt"
)

func main() {
	fmt.Println("Три числа.")
	var numberOne int
	var numberTwo int
	var numberThree int
	fmt.Println("Введите первое число:")
	fmt.Scan(&numberOne)
	fmt.Println("Введите второе число:")
	fmt.Scan(&numberTwo)
	fmt.Println("Введите третье число:")
	fmt.Scan(&numberThree)

	x := 0

	if numberOne >= 5 {
		x++
	}
	if numberTwo >= 5 {
		x++
	}
	if numberThree >= 5 {
		x++
	}
	if x > 0 {
		fmt.Println("Среди введенных чисел", x, "больше или равно 5")
	}
	if x == 0 {
		fmt.Println("Среди введенных чисел нет числа больше или равно 5")
	}
}
