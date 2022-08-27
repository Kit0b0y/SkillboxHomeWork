package main

import (
	"fmt"
)

func main() {
	var pointExam1 int
	var pointExam2 int
	var pointExam3 int
	sum := 275
	fmt.Println("Введите результат 1 экзамена:")
	fmt.Scan(&pointExam1)
	fmt.Println("Введите результат 2 экзамена:")
	fmt.Scan(&pointExam2)
	fmt.Println("Введите результат 3 экзамена:")
	fmt.Scan(&pointExam3)
	sumTotal := pointExam1 + pointExam2 + pointExam3
	fmt.Println("Сумма проходных баллов:", sum)
	fmt.Println("Количество набранных баллов:", sumTotal)
	if sumTotal >= sum {
		fmt.Println("Вы поступили")
	} else {
		fmt.Println("Вы не поступили")
	}
}
