package main

import "fmt"

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

	if numberOne > 5 {
		fmt.Println("Среди введенных чисел есть число больше 5")
	} else if numberTwo > 5 {
		fmt.Println("Среди введенных чисел есть число больше 5")
	} else if numberTwo > 5 {
		fmt.Println("Среди введенных чисел есть число больше 5")
	} else {
		fmt.Println("Среди введенных чисел нет числа больше 5")
	}
}
