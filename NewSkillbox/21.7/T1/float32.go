package main

import (
	"fmt"
	"math"
	//"math"
)

func formulaDifferentTypes(x int16, y uint8, z float32) float32 {
	var x1 float32 = float32(x)
	var y1 float32 = math.Pow(y, 2)
	return x1*2 + y1 + 3/z

}

func main() {
	fmt.Println(formulaDifferentTypes(1, 2, 3))

}
