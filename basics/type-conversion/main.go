package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Sprintf
	fmt.Println("fmt.Sprintf()......")
	var num1 = 1
	var num2 = 1.23
	var b = true
	var myChar byte = 'S'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("Type of str = %T, str = %q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("Type of str = %T, str = %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("Type of str = %T, str = %q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("Type of str = %T, str = %q\n", str, str)

	fmt.Println("=========================")

	// strconv package
	fmt.Println("strconv package......")
	var num3 = 1
	var num4 = 1.23
	var b2 = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("Type of str = %T, str = %q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("Type of str = %T, str = %q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("Type of str = %T, str = %q\n", str, str)

	fmt.Println("=========================")

	fmt.Println("Parse string type......")
	var str1 = "true"
	var b1 bool
	b1, _ = strconv.ParseBool(str1)
	fmt.Printf("Type of b1: %T, b1 = %v", b1, b1)

	var str2 = "1"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("Type of n1: %T, n1 = %v\n", n1, n1)
	fmt.Printf("Type of n2: %T, n2 = %v\n", n2, n2)

	var str3 = "1.23"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("Type of f1: %T, f1 = %v\n", f1, f1)

}
