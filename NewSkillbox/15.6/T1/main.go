package main

import "fmt"

func main() {

	var num [10]int

	for i := 0; i < 10; i++ {
		fmt.Print("Введите число:")
		fmt.Scan(&num[i])
	}
	evenCount := 0
	oddCount := 0

	for _, a := range num {
		if a%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	fmt.Println("================")
	fmt.Println("Всего четных чисел = ", evenCount)
	fmt.Println("Всего нечетных чисел  = ", oddCount)
}
