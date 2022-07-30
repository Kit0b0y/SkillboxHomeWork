package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ёлочка")
	fmt.Println("Введите высоту ёлки:")
	var height int
	fmt.Scan(&height)
	chars := 1 // начальное кол-во * в строке
	for i := height; i > 0; i-- {
		line := "*"                        // формируем строку состоящую из *
		indent := " "                      // добавляем отступы
		fmt.Printf("%s%s\n", indent, line) // выводим строку в консоль
		chars += 2                         // увеличиваем кол-во звёздочек на 2
	}
}
