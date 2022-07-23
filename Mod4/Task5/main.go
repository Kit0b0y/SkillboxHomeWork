package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ресторан.")

	var dayWeek int
	var numberGuest int
	var checkAmmount int

	fmt.Println("Введите день недели:")
	fmt.Scan(&dayWeek)
	fmt.Println("Введите число гостей:")
	fmt.Scan(&numberGuest)
	fmt.Println("Введите сумму чека:")
	fmt.Scan(&checkAmmount)

	var discont int
	var tips int
	var ammountPay int

	if dayWeek == 1 {
		discont = checkAmmount * 10 / 100
		fmt.Println("Скидка по понедельникам:", discont)
	}

	if dayWeek == 5 {
		if checkAmmount > 10000 {
			discont = checkAmmount * 5 / 100
			fmt.Println("Скидка по пятницам:", discont)
		}
	}
	if dayWeek != 1 {
	}
	if dayWeek != 5 {
	}
	if numberGuest > 5 {
		tips = checkAmmount * 10 / 100
		fmt.Println("Надбавка за обслуживание:", tips)
	} else {
		tips = 0
	}

	ammountPay = checkAmmount - discont + tips
	fmt.Println("Сумма к оплате:", ammountPay)
}
