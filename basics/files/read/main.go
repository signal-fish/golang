package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	workDir, _ := os.Getwd()
	path := workDir + "/hello.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Open file Error", err)
		return
	}

	const defaultBufSize = 4096
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("End of read file")
}
