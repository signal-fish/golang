package main

import "fmt"

func main() {
	var f1 float32 = 1.1
	var f2 = 2.2 // default float64
	var b = true
	var i1 int32 = 1
	var i2 int64 = 2
	i3 := 3
	var s1 = "Signal"
	s2 := '余'
	s3 := "幸浩"

	typeJudge(f1, f2, b, i1, i2, i3, s1, s2, s3)
}

func typeJudge(items ...interface{}) {
	fmt.Println("items =", items)
	for index, x := range items {
		switch x.(type) {
		case bool, float32, float64, int, int32, int64, string:
			fmt.Printf("Type of items[%v] is %T, items[%v] = %v\n", index, x, index, x)
		default:
			fmt.Printf("Unsure type")
		}

	}
}
