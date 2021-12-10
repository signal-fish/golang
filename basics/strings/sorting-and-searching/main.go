package main

import (
	"fmt"
	"sort"
)

func main() {
	countries := sort.StringSlice{"China", "Japan", "Ukraine", "America"}
	fmt.Println("Before sorting, countries =", countries)
	countries.Sort()
	fmt.Println("After sorting, countries =", countries)
	n := countries.Search("Ukraine")
	fmt.Println("Result: ", n, countries[n])
}
