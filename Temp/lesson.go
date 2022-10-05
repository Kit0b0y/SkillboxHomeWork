package main

import "fmt"

const oRows = 3
const oCols = 5
const nCols = 4

func multiMat(m1 [oRows][oCols]int64, m2 [oCols][nCols]int64) [oRows][nCols]int64 {
	var m [oRows][nCols]int64
	for i := 0; i < oRows; i++ {
		for j := 0; j < nCols; j++ {
			for k := 0; k < oCols; k++ {
				m[i][j] = m[i][j] + m1[i][k]*m2[k][j]
			}
		}
	}
	return m
}

func main() {

	m2 := [oCols][nCols]int64{
		{10, 20, 30, 40},
		{11, 22, 33, 44},
		{55, 66, 77, 88},
		{21, 32, 43, 54},
		{16, 15, 32, 45},
	}
	m1 := [oRows][oCols]int64{
		{5, 4, 6, 3, 6},
		{54, 56, 43, 89, 58},
		{34, 81, 72, 83, 74},
	}
	fmt.Println(multiMat(m1, m2))
}
