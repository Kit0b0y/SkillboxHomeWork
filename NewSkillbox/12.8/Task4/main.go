package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	i := 1
	var b bytes.Buffer
	var fileName = "file.txt"

	for {
		var text string

		fmt.Print("Введите, пожалуйста, текст: ")
		fmt.Scan(&text)

		if strings.ToLower(text) == "exit" {
			fmt.Print("Выход из программы...\n")
			break
		} else {
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			b.WriteString(strconv.Itoa(i) + ". " + currentTime + " " + text + "\n")
		}
		i++
	}

	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		fmt.Println("Ошибка записи в файл: ", err)
	}
}
