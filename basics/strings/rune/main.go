package main

import "fmt"

/*
1. A rune is a builtin type in Go and it's the alias of int32.
2. Rune represents a Unicode code point in Go.
*/

func main() {
	name := "よこうこう"
	fmt.Printf("String: %s\n", name)
	printChars(name)

	fmt.Println()

	emoji := "💘"
	fmt.Printf("String: %s\n", emoji)
	printChars(emoji)
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for _, rune := range s {
		fmt.Printf("%c ", rune)
	}
}
