package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "Go is an Open source programming Language that makes it Easy\n to build simple, reliable, and efficient Software"
	fmt.Println("Определение количества слов, начинающихся с большой буквы в строке:\n", str)
	count := 0
	for _, i := range str {
		if unicode.IsUpper(i) {
			count++
		}
	}
	fmt.Println("Строка содержит", count, "слов с большой буквы")
}
