package main

import (
	"fmt"
)

var coin1 int
var coin2 int
var coin3 int
var sum int
var rest int // остаток суммы для внесения монет
var ext int  // сумма для сдачи

func main() {
	fmt.Println("Банкомат.\n")
	fmt.Println("Введите сумму:")
	fmt.Scan(&sum)
	fmt.Println("Опустите монеты суммой номинала", sum)
	fmt.Scan(&coin1)

	if coin1 > sum {
		ext = coin1 - sum
		fmt.Println("Ваша сдача составит:", ext)
		return
	}
	if coin1 == sum {
		fmt.Println("Вы ввели нужную сумму")
		return
	}
	if coin1 < sum {
		rest = sum - coin1
		fmt.Println("Осталось внести", rest)
	}
	fmt.Scan(&coin2)
	if coin2 > rest {
		ext = coin2 - rest
		fmt.Println("Ваша сдача составит:", ext)
		return
	}

	if coin2 == rest {
		fmt.Println("Вы ввели нужную сумму")
		return
	}
	if coin2 < rest {
		rest = sum - coin1 - coin2
		fmt.Println("Для оплаты без сдачи внесите 1 монету номиналом", rest)
	}
	fmt.Scan(&coin3)

	if coin3 == rest {
		fmt.Println("Вы внесли нужную сумму без сдачи")
	}
	if coin3 > rest {
		ext = (coin1 + coin2 + coin3) - sum
		fmt.Println("Ваша сдача составит:", ext)
	}
	if coin3 < rest {
		fmt.Println("Операция не возможна. Недостаток внесенных средств.")
	}
}
