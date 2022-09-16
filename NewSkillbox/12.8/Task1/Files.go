package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	for {
		var text string
		fmt.Print("Введите текст: ")
		fmt.Scan(&text)

		if text == "exit" {
			fmt.Print("Выход из программы...\n")
			break
		} else {
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			file.WriteString(currentTime + " " + text + "\n")
		}
	}
}
