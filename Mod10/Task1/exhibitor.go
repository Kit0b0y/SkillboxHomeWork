package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Введите значение х:")
	var x int
	fmt.Scan(&x)
	fmt.Println("Введите значение точности вычисления:")
	var y int
	fmt.Scan(&y)
	var epsilon, fact float64
	fact = 1
	epsilon = 1 / math.Pow(10, float64(y))
	n := 1
	var ex float64 = 1
	var preEx float64 = 1

	for {
		fact *= float64(n)
		ex += math.Pow(float64(x), float64(n)) / fact
		if math.Abs(ex-preEx) < epsilon {
			fmt.Println("Значение экспоненты:", ex)
			break
		}
		n++
		preEx = ex
	}
}
