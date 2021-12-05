package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1.len()
	str1 := "Signal余"
	fmt.Println("str1 =", str1)
	fmt.Println("len(str1) =", len(str1))

	fmt.Println("=========================")

	// []rune()
	str2 := "Signal余"
	r := []rune(str2)
	fmt.Println("str2 =", str2)
	fmt.Print("Traverse str2: ")
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c ", r[i])
	}
	fmt.Println()

	fmt.Println("=========================")

	// strconv.Atoi()
	str3, err := strconv.Atoi("1")
	if err != nil {
		fmt.Println("Error occurred", err)
	} else {
		fmt.Println("str3 =", str3)
	}

	fmt.Println("=========================")

	// strconv.Itoa()
	str4 := strconv.Itoa(1)
	fmt.Printf("Type of str4 = %T, str4 = %v\n", str4, str4)

	fmt.Println("=========================")

	// []byte
	var bytes = []byte("Signal余")
	fmt.Println("bytes =", bytes)

	fmt.Println("=========================")

	// string()
	str5 := string([]byte{228, 189, 153})
	fmt.Println("str5 =", str5)

	fmt.Println("=========================")

	// strconv.FormatInt()
	b99 := strconv.FormatInt(99, 2)
	fmt.Println("b99 =", b99)
	o99 := strconv.FormatInt(99, 8)
	fmt.Println("o99 =", o99)
	h99 := strconv.FormatInt(99, 16)
	fmt.Println("h99 =", h99)

	fmt.Println("=========================")

	// strings.Contains()
	res := strings.Contains("signal", "sig")
	fmt.Println(`strings.Contains("signal") =`, res)

	fmt.Println("=========================")

	// strings.Count()
	count := strings.Count("signal fish", "i")
	fmt.Println(`strings.Count("signal fish", "i") =`, count)

	fmt.Println("=========================")

	// strings.EqualFold()
	equal := strings.EqualFold("PI", "pi")
	fmt.Println(`strings.EqualFold("PI", "pi") ==>`, equal)
	fmt.Println(`"PI" == "pi" ==>`, "PI" == "pi")

	fmt.Println("=========================")

	// strings.Index(), strings.LastIndex()
	index := strings.Index("_yu_yu_", "yu")
	fmt.Println(`strings.Index("_yu_yu_", "yu") ==>`, index)
	lastIndex := strings.LastIndex("_yu_yu_", "yu")
	fmt.Println(`strings.LastIndex("_yu_yu_", "yu") ==>`, lastIndex)

	fmt.Println("=========================")

	// strings.Replace()
	str6 := "Signal Yu Yu Yu"
	fmt.Println("str6 =", str6)
	str6_1 := strings.Replace(str6, "Yu", "Fish", 1)
	fmt.Println("str6_1 =", str6_1)
	str6_2 := strings.Replace(str6, "Yu", "Fish", 2)
	fmt.Println("str6_2 =", str6_2)
	str6_3 := strings.Replace(str6, "Yu", "Fish", -1)
	fmt.Println("str6_3 =", str6_3)

	fmt.Println("=========================")

	// strings.Split()
	strArr := strings.Split("signal,fish,yu,marsha", ",")
	fmt.Println("strArr =", strArr)

	fmt.Println("=========================")

	// strings.ToLower(), strings.ToUpper()
	str7 := "SigNAL"
	fmt.Println("str7 =", str7)
	str7_1 := strings.ToLower(str7)
	fmt.Println("str7_1 =", str7_1)
	str7_2 := strings.ToUpper(str7)
	fmt.Println("str7_2 =", str7_2)

	fmt.Println("=========================")

	// Trim
	str8 := "  **fish **"
	fmt.Println("str8 =", str8)
	str8_1 := strings.TrimSpace(str8)
	fmt.Println("str8_1 =", str8_1)
	str8_2 := strings.Trim(str8, "*")
	fmt.Println("str8_2 =", str8_2)
	str8_3 := strings.TrimLeft(str8, " ")
	fmt.Println("str8_3 =", str8_3)
	str8_4 := strings.TrimRight(str8, "*")
	fmt.Println("str8_4 =", str8_4)

	fmt.Println("=========================")

	// strings.HasPrefix()
	str9 := "ftp://192.168.10.1"
	ftp := strings.HasPrefix(str9, "ftp")
	fmt.Println("ftp =", ftp)

	fmt.Println("=========================")

	// strings.HasSuffix()
	str10 := "signalfish.jpg"
	jpeg := strings.HasSuffix(str10, "jpeg")
	fmt.Println("jpeg =", jpeg)
}
