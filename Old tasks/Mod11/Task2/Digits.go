package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "a10 10 20b 20 30c30 30 dd"
	fmt.Println("Исходная строка:\n", s)
	space := " "
	for strings.Contains(s, space) {
		spaceIndex := strings.Index(s, space)
		w := s[:spaceIndex]
		i, err := strconv.Atoi(w)
		n := 0
		s = s[spaceIndex+1:]
		s = strings.Trim(s, space)

		if err == nil {
			n = i
			fmt.Println("Десятичное число:", n)
		}
	}
}
