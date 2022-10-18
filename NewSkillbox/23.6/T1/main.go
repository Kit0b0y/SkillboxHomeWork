package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var oddNumbers, evenNumbers []int

	oddNumbers, evenNumbers = splitOddAndEvenNumbers(numbers)

	fmt.Println(oddNumbers)
	fmt.Println(evenNumbers)
}

func splitOddAndEvenNumbers(numbers []int) (oddNumbers, evenNumbers []int) {
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			evenNumbers = append(evenNumbers, i)
		} else {
			oddNumbers = append(oddNumbers, i)
		}
	}
	return
}
