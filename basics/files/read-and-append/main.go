package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	workDir, _ := os.Getwd()
	path := workDir + "/test.txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str := "Golang is awesome\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println("Write String Error:", err)
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("Flush Error")
	}
}
