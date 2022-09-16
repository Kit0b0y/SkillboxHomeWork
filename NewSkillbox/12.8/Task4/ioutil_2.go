package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//var b bytes.Buffer

	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

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

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}
