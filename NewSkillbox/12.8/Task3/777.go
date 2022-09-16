package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Create("test.txt")
	if err := os.Chmod("test.txt", 0111); err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("Someone text")
	if err != nil {
		fmt.Println("Can`t write file, ", err)
		return
	}
	openFile, err := os.Open(f.Name())
	if err != nil {
		panic(err)
		return
	}
	data := make([]byte, 64)
	for {
		_, err := openFile.Read(data)
		if err == io.EOF {
			break
		}
	}
	fmt.Println(string(data))

}
