package main

import "math"

func main() {

	var n int = 1
	fact := 1
	var x float64 = 1
	var ex float64
	var pex float64 = 0
	epsilon := 0.1

	for {
		for i := 1; i <= n; i++ {
			fact *= i
		}
		ex += math.Pow(x, float64(n)) / float64(fact)
		if math.Abs(ex-pex) < epsilon {
			println(ex)
			break
		}
		n++
		pex = ex
	}
}
