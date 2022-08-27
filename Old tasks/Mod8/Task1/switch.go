package main

import "fmt"

func main() {
	fmt.Println("Времена года")
	var word string
	fmt.Println("Введите месяц:")
	_, _ = fmt.Scan(&word)
	switch word {
	case "декабрь", "январь", "февраль":
		fmt.Println("Время года - Зима")
	case "март", "апрель", "май":
		fmt.Println("Время года - Весна")
	case "июнь", "июль", "август":
		fmt.Println("Время года - Лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("Время года - Осень")
	default:
		fmt.Println("Нет такого времени года")
	}
}
