package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	workDir, _ := os.Getwd()
	src := workDir + "/test.txt"
	dest := workDir + "/copy.txt"
	_, err := copyFile(dest, src)

	if err == nil {
		fmt.Println("Copy Completed")
	} else {
		fmt.Println("Copy Error Occurred", err)
	}
}

func copyFile(dest, src string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Open File Error:", err)
		return 0, err
	}
	reader := bufio.NewReader(srcFile)

	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Open File Error", err)
		return 0, err
	}

	writer := bufio.NewWriter(destFile)

	return io.Copy(writer, reader)
}
