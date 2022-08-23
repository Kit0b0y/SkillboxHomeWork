package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "Go lang"
	s2 := "go lang"
	count := 0
	a1 := strings.Split(s1, " ")
	a2 := strings.Split(s2, " ")
	fmt.Println(a1)
	fmt.Println(a2)
	space := " "
	spaceIndex := strings.Index(s1, space)
	w1 := a1[spaceIndex:]
	w2 := a2[spaceIndex:]
	if strings.Contains(w1, w2) {
		count++
		fmt.Println(count)
	}
}
