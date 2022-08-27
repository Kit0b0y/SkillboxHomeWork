package main

import (
	"fmt"
	"math"
)

var a, b, c float64
var x1, x2 float64 //Корени уравнения
var d float64      //Дискриминант

func main() {
	fmt.Println("Квадратное уравнение.")
	fmt.Println("Введите первый коэффициент:")
	fmt.Scan(&a)
	fmt.Println("Введите второй коэффициент:")
	fmt.Scan(&b)
	fmt.Println("Введите свободную переменную:")
	fmt.Scan(&c)

	d = math.Pow(b, 2) - 4*a*c
	//fmt.Println(d)

	if d > 0 {
		x1 = (-b + math.Sqrt(d)) / 2 * a
		x2 = (-b - math.Sqrt(d)) / 2 * a
		fmt.Println("Результат:", x1, x2)
	}
	if d == 0 {
		x1 = (-b / 2 * a)
		fmt.Println("Результат:", x1)
	}
	if d < 0 {
		fmt.Println("Уравнение не имеет корней:")
	}
}
