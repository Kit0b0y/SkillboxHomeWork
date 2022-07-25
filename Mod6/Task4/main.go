package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0
	c := 0

	for a < 10 {
		a = a + 1
		fmt.Println("a =:", a)
		if a > 10 {
			continue
		}
		for b < 100 {
			b = b + 1
			fmt.Println("b =:", b)
			if b > 100 {
				continue
			}
			for c < 1000 {
				c = c + 1
				fmt.Println("c =:", c)
				if c > 1000 {
					break
				}
			}
		}
	}
}
