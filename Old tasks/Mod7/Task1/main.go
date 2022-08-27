package main

import (
	"fmt"
)

func main() {
	var number, count int
	for i := 99999; i <= 999999; i++ {
		number = i
		if number > 99999 && number <= 999999 {
			num1 := (number - (number % 100000)) / 100000
			num2 := (number%100000 - (number % 10000)) / 10000
			num3 := (number%10000 - (number % 1000)) / 1000
			num4 := (number%1000 - (number % 100)) / 100
			num5 := (number%100 - (number % 10)) / 10
			num6 := number % 10
			if num1 == num6 && num2 == num5 && num3 == num4 {
				count++
			}
		}
	}
	fmt.Println(count)
}
