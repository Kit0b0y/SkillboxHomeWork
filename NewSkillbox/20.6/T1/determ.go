package main

import "fmt"

func main() {

	var d [3][3]int

	fmt.Println("Введите числа матрицы 3х3:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&d[i][j])
		}
	}

	x := (d[1][1] * d[2][2]) - (d[2][1] * d[1][2])
	y := (d[1][0] * d[2][2]) - (d[2][0] * d[1][2])
	z := (d[1][0] * d[2][1]) - (d[2][0] * d[1][1])

	determinant := (d[0][0] * x) - (d[0][1] * y) + (d[0][2] * z)
	fmt.Println("Определитель матрицы равен:", determinant)
}
