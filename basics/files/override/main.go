package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	workDir, _ := os.Getwd()
	path := workDir + "/test.txt"
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file error = %v\n", err)
		return
	}

	str := "Golang is awesome\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println("Write String Error")
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Flush Error")
	}
}
