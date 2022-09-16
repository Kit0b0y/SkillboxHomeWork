package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {

	file, err := os.Create("Log.txt")
	if err != nil {
		fmt.Println("Не смогли создать файл", err)
		return
	}
	defer file.Close()
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	fmt.Println("Введите число от 1 до 100")
	file.WriteString("Введите число от 1 до 100 \n")
	for {
		var answer int
		for {
			_, _ = fmt.Scan(&answer)
			file.WriteString(fmt.Sprintf("Введите число %d\n", answer))
			if answer < 1 || answer > 100 {
				fmt.Println("Число должно быть в диапазоне от 1 до 100")
				file.WriteString("Число должно быть в диапазоне от 1 до 100 \n")
			} else {
				break
			}
		}
		if answer == n {
			fmt.Println("Ура! Число угадано")
			file.WriteString("Ура! Число угадано \n")
			return
		} else if answer < n {
			fmt.Println("Загаданное число больше")
			file.WriteString("Загаданное число больше \n")
		} else {
			fmt.Println("Загаданное число меньше")
			file.WriteString("Загаданное число меньше \n")
		}
	}
	f, err := os.Open("Log.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 256)
	if _, err := io.ReadFull(f, buf); err != nil {
		panic(err)
	}
	fmt.Println("=======")
	fmt.Printf("%s\n", buf)
}
