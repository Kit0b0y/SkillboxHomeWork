package main

import "fmt"

func printFuncResult(f func(int, int) int, a int, b int, oper string) {
	defer fmt.Println(f(a, b))

	var operationString string
	switch oper {
	case "+":
		operationString = "сложение"
	case "-":
		operationString = "вычитание"
	case "*":
		operationString = "умножение"
	}

	fmt.Printf("Результат операции <%s>: ", operationString)
}

func main() {
	printFuncResult(func(a, b int) int {
		return a + b
	}, 10, 20, "+")
	printFuncResult(func(a, b int) int {
		return a * b
	}, 10, 20, "*")
	printFuncResult(func(a, b int) int {
		return a - b
	}, 10, 20, "-")
}
