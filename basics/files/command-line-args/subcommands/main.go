package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addcmd := flag.NewFlagSet("add", flag.ExitOnError)
	aAdd := addcmd.Int("a", 0, "The value of a")
	bAdd := addcmd.Int("b", 0, "The value of b")
	err := addcmd.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("Parse Error:", err)
	}

	mulcmd := flag.NewFlagSet("mul", flag.ExitOnError)
	aMul := mulcmd.Int("a", 0, "The value of a")
	bMul := mulcmd.Int("b", 0, "The value of b")
	err = mulcmd.Parse(os.Args[2:])
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
