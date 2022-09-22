package main

import "fmt"

var evenCount, oddCount int

func countEvenNums(evoddarr []int) int {
	evenCount = 0
	fmt.Print("nList of Even Numbers  = ")
	for _, a := range evoddarr {
		if a%2 == 0 {
			fmt.Println(a, " ")
			evenCount++
		}
	}
	return evenCount
}
func countOddNums(evoddarr []int) int {
	oddCount = 0
	fmt.Print("nList of Odd Numbers   = ")
	for _, a := range evoddarr {
		if a%2 != 0 {
			fmt.Println(a, " ")
			oddCount++
		}
	}
	return oddCount
}
func main() {
	var size int
	fmt.Print("Enter the Even Odd Array Size = ")
	fmt.Scan(&size)
	evoddarr := make([]int, size)
	fmt.Print("Enter the Even Odd Array Items  = ")
	for i := 0; i < size; i++ {
		fmt.Scan(&evoddarr[i])
	}
	evenCount = countEvenNums(evoddarr)
	oddCount = countOddNums(evoddarr)
	fmt.Println("nThe Total Number of Even Numbers = ", evenCount)
	fmt.Println("The Total Number of Odd Numbers  = ", oddCount)
}
