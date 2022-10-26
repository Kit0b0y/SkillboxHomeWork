package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func main() {
	var array = [size]int{}
	var generate string

Loop:
	for {
		fmt.Print("Хотите самостоятельно заполнить массив (Y/N)? ")
		fmt.Scan(&generate)

		switch generate {
		case "N":
			array = generateArray(size)
			break Loop
		case "Y":
			for i := 0; i < size; i++ {
				fmt.Printf("Пожалуйста, введите, элемент массива [%v]: ", i)
				fmt.Scan(&array[i])
			}
			break Loop
		default:
			fmt.Print("Пожалуйста, повторите ввод. Не распознал ваш ответ :(\n")
		}
	}

	fmt.Printf("Исходный массив: %v\n", array)

	array = insertionSort(array)

	fmt.Print("Отсортированный массив: ", array)
}

func generateArray(arraySize int) (array [size]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < arraySize; i++ {
		array[i] = rand.Intn(50) - rand.Intn(50)
	}
	return
}

func insertionSort(items [size]int) [size]int {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	return items
}
