package main

import (
	"fmt"
	"math"
)

func main() {
	var count8 int = 0
	var count16 int = 0
	for i := 0; i <= math.MaxUint32; i++ {
		if uint8(i) == 0 {
			count8++
		}
		if uint16(i) == 0 {
			count16++
		}
	}
	fmt.Println("Количество переполнений uint8:", count8)
	fmt.Println("Количество переполнений uint16:", count16)
}
