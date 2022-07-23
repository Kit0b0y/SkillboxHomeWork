package main

import (
	"fmt"
)

func main() {
	fmt.Println("Банкомат.")
	var bn int
	fmt.Println("Введите сумму снятия со счета:")
	fmt.Scan(&bn)

	if bn > 100000 {
		fmt.Println("Операция не возможна, Вы превысили сумму.")
	}
	if bn <= 100000 {
		if bn%100 == 0 {
			fmt.Println("Операция успешно выполнена.\nВы сняли", bn, "рублей.")

		} else if bn%100 != 0 {
			fmt.Println("Операция не возможна, введите сумму, кратную 100 рублям.")
		}
	}
}
