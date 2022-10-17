package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int
	array := []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array)

	fmt.Print("Введите число, первое вхождение которого хотите найти: ")
	fmt.Scan(&number)

	index, err := indexOf(array, number)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Первое вхождение числа - по индексу: %v\n", index)
	}
}

func indexOf(array []int, key int) (int, error) {
	lo := 0
	hi := len(array) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < array[mid] {
			hi = mid - 1
		} else if key > array[mid] {
			lo = mid + 1
		} else if mid == 0 || key != array[mid-1] {
			return mid, nil
		} else {
			//нашли дубликат
			hi = mid - 1
		}
	}

	return -1, errors.New("Элемент не найден!")
}
