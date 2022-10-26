package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func main() {

	var generate string
	var array [size]int

Loop:
	for {
		fmt.Print("Хотите самостоятельно заполнить массив (Y/N)? ")
		fmt.Scan(&generate)

		switch generate {
		case "N":
			array = generateArray(size)
			break Loop
		case "Y":
			{
				for i := 0; i < size; i++ {
					fmt.Printf("Пожалуйста, введите, элемент массива [%v]: ", i)
					fmt.Scan(&array[i])
				}
				break Loop
			}

		default:
			fmt.Print("Пожалуйста, повторите ввод. Не распознал ваш ответ :(\n")
		}
	}

	fmt.Printf("Исходный массив: %v\n", array)

	reverseBubbleSort := func(items [size]int) [size]int {
		for i := 0; i < size; i++ {
			for j := size - 1; j >= i+1; j-- {
				if items[j] > items[j-1] {
					items[j], items[j-1] = items[j-1], items[j]
				}
			}
		}
		return items
	}

	fmt.Printf("Отсортированный массив: %v", reverseBubbleSort(array))
}

func generateArray(arraySize int) [size]int {
	var items [size]int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < arraySize; i++ {
		items[i] = rand.Intn(50) - rand.Intn(50)
	}
	return items
}
