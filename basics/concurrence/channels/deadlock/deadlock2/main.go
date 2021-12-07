package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "Signal"
	ch <- "Fish"
	ch <- "Neupoeyu"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
