package main

import "fmt"

func main() {
	fmt.Println("Шахматная доска.")
	fmt.Println("Введите ширину:")
	var line int
	fmt.Scan(&line)
	fmt.Println("Введите высоту:")
	var column int
	fmt.Scan(&column)
	//l:
	for i := 0; i < column; i++ {
		for j := 0; j < line; j++ {
			if i%2 == 0 && j%2 == 0 || i%2 != 0 && j%2 != 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
