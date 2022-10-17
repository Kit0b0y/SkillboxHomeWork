package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const count = 15

func main() {
	var number int
	array := generateArray()
	fmt.Println(array)

	fmt.Print("Пожалуйста, введите число, которое нужно найти в массиве: ")
	fmt.Scan(&number)

	arrayElementsLeft, err := arrayElementsAfterIndex(array, number)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Число найдено в массиве. Элементов после числа: %v\n", arrayElementsLeft)
	}
}

func generateArray() (array []int) {
	array = make([]int, count)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		array[i] = rand.Intn(20)
	}
	return
}
func arrayElementsAfterIndex(array []int, number int) (int, error) {
	index, arrayElementsLeft := -1, -1

	for i := 0; i < count; i++ {
		if array[i] == number {
			index = i
			arrayElementsLeft = len(array[index:]) - 1
		}
	}
	if index == -1 {
		return arrayElementsLeft, errors.New("Число не найдено в массиве.")
	}
	return arrayElementsLeft, nil
}
