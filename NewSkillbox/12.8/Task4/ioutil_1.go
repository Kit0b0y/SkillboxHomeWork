package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {
	var b bytes.Buffer

	for {
		var text string
		fmt.Print("Введите текст: ")
		fmt.Scan(&text)

		if text == "exit" {
			fmt.Print("Выход из программы...\n")
			break
		} else {
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			b.WriteString(currentTime + " " + text + "\n")
		}
	}
	fileName := "file.txt"
	if err := os.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

}
