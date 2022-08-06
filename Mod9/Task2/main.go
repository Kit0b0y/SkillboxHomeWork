package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int16
	fmt.Scan(&a)
	fmt.Scan(&b)
	var res = int64(a) * int64(b)
	if res > 0 {
		if res <= math.MaxUint8 {
			fmt.Println("uint8")
		} else if res <= math.MaxUint16 {
			fmt.Println("uint16")
		} else {
			fmt.Println("uint32")
		}
	} else {
		if res >= math.MinInt8 {
			fmt.Println("int8")
		} else if res >= math.MinInt16 {
			fmt.Println("int16")
		} else {
			fmt.Println("int32")
		}
	}
}
