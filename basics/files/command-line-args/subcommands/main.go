package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	aAdd := addCmd.Int("a", 0, "The value of a")
	bAdd := addCmd.Int("b", 0, "The value of b")
	err := addCmd.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("Parse Error:", err)
	}

	mulCmd := flag.NewFlagSet("mul", flag.ExitOnError)
	aMul := mulCmd.Int("a", 0, "The value of a")
	bMul := mulCmd.Int("b", 0, "The value of b")
	err = mulCmd.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("Parse Error:", err)
	}

	switch os.Args[1] {
	case "add":
		fmt.Println(*aAdd + *bAdd)
	case "mul":
		fmt.Println((*aMul) * (*bMul))
	default:
		fmt.Println("expected add or mul command")
		os.Exit(1)
	}
}
