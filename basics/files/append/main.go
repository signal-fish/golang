package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	workDir, _ := os.Getwd()
	path := workDir + "/test.txt"
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open file error:", err)
		return
	}

	str := "Cute Gopher\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			fmt.Println("Write String Error:", err)
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("Flush Error:", err)
	}
}
