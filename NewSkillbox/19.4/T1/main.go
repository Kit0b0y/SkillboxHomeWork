package main

import "fmt"

func main() {
	var arrThree [9]int
	arrOne := [4]int{1, 2, 3, 4}
	arrTwo := [5]int{5, 6, 7, 8, 9}
	s := arrThree[:0]
	s = append(s, arrOne[:]...)
	s = append(s, arrTwo[:]...)

	fmt.Println(arrThree)
}
