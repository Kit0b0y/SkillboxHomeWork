package main

import (
	"fmt"
	"math/rand"
	"time"
)

func firstFunc(x1, y1, x2, y2, x3, y3 int) {

	fmt.Printf("Исходные значения: (%v, %v); (%v, %v); (%v, %v)\n", x1, y1, x2, y2, x3, y3)
}

func newPoint(x, y *int) {
	*x = 2*(*x) + 10
	*y = -3*(*y) - 5

}

func getRandom(n1, n2 int) (int, int) {
	return rand.Intn(n1), rand.Intn(n2)

}

func main() {
	rand.Seed(time.Now().UnixNano())
	x1, y1 := getRandom(10, 20)
	x2, y2 := getRandom(10, 20)
	x3, y3 := getRandom(10, 20)

	firstFunc(x1, y1, x2, y2, x3, y3)

	newPoint(&x1, &y1)
	newPoint(&x2, &y2)
	newPoint(&x3, &y3)
	fmt.Printf("Новые значения: (%v, %v); (%v, %v); (%v, %v)\n", x1, y1, x2, y2, x3, y3)

}
