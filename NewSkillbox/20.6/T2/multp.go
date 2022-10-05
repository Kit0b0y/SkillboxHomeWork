package main

import "fmt"

const rw = 3
const cl1 = 5
const cl2 = 4

func multiMat(m1 [rw][cl1]int64, m2 [cl1][cl2]int64) [rw][cl2]int64 {
	var mat [rw][cl2]int64
	for i := 0; i < rw; i++ {
		for j := 0; j < cl2; j++ {
			for k := 0; k < cl1; k++ {
				mat[i][j] = mat[i][j] + m1[i][k]*m2[k][j]
			}
		}
	}
	return mat
}

func main() {

	m1 := [rw][cl1]int64{
		{5, 4, 6, 3, 6},
		{54, 56, 43, 89, 58},
		{34, 81, 72, 83, 74},
	}
	m2 := [cl1][cl2]int64{
		{10, 20, 30, 40},
		{11, 22, 33, 44},
		{55, 66, 77, 88},
		{21, 32, 43, 54},
		{16, 15, 32, 45},
	}
	fmt.Println("Результат:")
	fmt.Println(multiMat(m1, m2))
}
