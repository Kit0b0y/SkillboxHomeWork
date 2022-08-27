package main

import "fmt"

func main() {
	fmt.Println("Дни недели")
	var workDay string
	fmt.Println("Введите будний день недели: пн, вт, ср, чт, пт:")
	_, _ = fmt.Scan(&workDay)
	switch {
	case workDay == "пн":
		fmt.Println("понедельник")
		fallthrough
	case workDay == "вт":
		fmt.Println("вторник")
		fallthrough
	case workDay == "ср":
		fmt.Println("среда")
		fallthrough
	case workDay == "чт":
		fmt.Println("четверг")
		fallthrough
	case workDay == "пт":
		fmt.Println("пятница")
	default:
		fmt.Println("Не правильное значение ввода")
	}
}
