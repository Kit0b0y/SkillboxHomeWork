package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var l = len(arr)
	newArr := make([]int, l)
	for i, v := range arr {
		_ = append(newArr[:l-1-i], v)
	}

	fmt.Println(newArr)
}
