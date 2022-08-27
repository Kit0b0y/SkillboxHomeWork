package main

import "fmt"

func main() {

	var price, discount, rate, rateLim int

	fmt.Println("Размер скидки")
	fmt.Println("Введите цену товара:")
	fmt.Scan(&price)
	fmt.Println("Введите размер скидки:")
	fmt.Scan(&discount)
	rate = price * discount / 100
	rateLim = price * 30 / 100

	if discount <= 30 && rate <= 2000 {
		fmt.Println("Размер вашей скидки составит:", rate)
	} else if rate > 2000 {
		fmt.Println("Ваша скидка составит: 2000")
	} else if discount > 30 {
		fmt.Println("Ваша скидка составит:", rateLim)
	}
}
