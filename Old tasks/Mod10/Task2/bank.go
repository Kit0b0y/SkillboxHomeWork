package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Введите сумму вклада:")
	var sum float64
	fmt.Scan(&sum)
	fmt.Println("Введите процент капитализации:")
	var percent float64
	fmt.Scan(&percent)
	fmt.Println("Введите срок вклада (в годах):")
	var term int
	fmt.Scan(&term)
	var profitBank float64 = 0
	var result float64

	for i := 0; i < term*12; i++ {
		result = sum + sum*percent/100
		sum = math.Trunc(result*100) / 100
		profitBank += result - sum
	}
	fmt.Println("Сумма вкладаК с процентами:", sum)
	fmt.Println("Выгода банка:", profitBank)
}
