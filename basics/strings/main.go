package main

import "fmt"

// A string is a slice of bytes in Go.

func main() {
	name := "Fish"
	fmt.Printf("String: %s\n", name)
	printBytes(name)

	fmt.Println()

	emoji := "💘"
	fmt.Printf("String: %s\n", emoji)
	printBytes(emoji)
}

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}
