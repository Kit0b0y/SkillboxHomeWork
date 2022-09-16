package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	for {
		inf, err := file.Stat()
		if err != nil {
			panic(err)
		}
		if inf.Size() == 0 {
			fmt.Println("File is empty")
		}
		break
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}
}
