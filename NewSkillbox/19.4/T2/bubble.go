package main

import "fmt"

func main() {
	sample := []int{3, 4, 5, 2, 6, 1}
	fmt.Println("\nИсходный массив:")
	fmt.Println(sample)

	bubbleSort(sample)

}

func bubbleSort(arr []int) {
	var lenghtArr int = len(arr)
	for i := 0; i < lenghtArr-1; i++ {
		for j := 0; j < lenghtArr-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("\nМассив после сортировки:")
	fmt.Println(arr)

}
