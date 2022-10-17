package main

import (
	"fmt"
	"math"
	//"math"
)

func formulaDifferentTypes(x int16, y uint8, z float32) float32 {
	return 2.0*float32(x) + float32(math.Pow(float64(y), 2.0)) - (3 / float32(z))
}

func main() {
	var (
		x int16
		y uint8
		z float32
	)

	for {
		fmt.Print("Введите x: ")
		fmt.Scan(&x)
		if x < math.MinInt16 || x > math.MaxInt16 {
			fmt.Print("Переполнение типа Int16. Попробуйте снова!\n")
		} else {
			break
		}
	}

	for {
		fmt.Print("Введите y: ")
		fmt.Scan(&y)
		if y < 0 || y > math.MaxUint8 {
			fmt.Print("Переполнение типа UInt8. Попробуйте снова!\n")
		} else {
			break
		}
	}

	for {
		fmt.Print("Введите z: ")
		fmt.Scan(&z)
		if z > math.MaxFloat32 {
			fmt.Print("Переполнение типа Float32. Попробуйте снова!\n")
		} else {
			break
		}
	}

	fmt.Printf("S = 2*x + y^2 - 3/z = %v", formulaDifferentTypes(x, y, z))
}
