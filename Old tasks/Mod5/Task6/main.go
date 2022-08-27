package main

import (
	"fmt"
)

var ticket int
var number1, number2, number3, number4 int
var digit1, digit2, digit3, digit4 int

func main() {
	fmt.Println("Счастливый билет")
	fmt.Println("Введите номер билета:")
	fmt.Scan(&ticket)

	number4 = ticket / 10
	//fmt.Println(number4)
	digit4 = ticket % 10
	//fmt.Println(digit4)

	number3 = number4 / 10
	//fmt.Println(number3)
	digit3 = number4 % 10
	//fmt.Println(digit3)

	number2 = number3 / 10
	//fmt.Println(number2)
	digit2 = number3 % 10
	//fmt.Println(digit2)

	number1 = number2 / 10
	//fmt.Println(number1)
	digit1 = number2 % 10
	//fmt.Println(digit1)

	if digit1 == digit4 && digit2 == digit3 {
		fmt.Println("Вам выпал зеркальный билет!")
		return
	}
	if (digit1+digit2 == digit3+digit4) != (digit1 == digit4 && digit2 == digit3) {
		fmt.Println("Вам выпал счастливый билет!")
	} else {
		fmt.Println("Обычный билет")
	}
}
