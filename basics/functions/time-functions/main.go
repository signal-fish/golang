package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now =", now)
	fmt.Println("now.Year =", now.Year())
	fmt.Println("now.Month =", now.Month())
	fmt.Println("now.Month =", int(now.Month()))
	fmt.Println("now.Day =", now.Day())
	fmt.Println("now.Hour =", now.Hour())
	fmt.Println("now.Minute =", now.Minute())
	fmt.Println("now.Second =", now.Second())
	fmt.Println("now.Unix() =", now.Unix())
	fmt.Println("now.UnixNano() =", now.UnixNano())
	fmt.Println("time.NanoSecond =", time.Nanosecond)
	fmt.Println("time.Microsecond =", time.Microsecond)
	fmt.Println("time.Millisecond =", time.Millisecond)
	fmt.Println("time.Second =", time.Second)
	fmt.Println("time.Minute =", time.Minute)
	fmt.Println("time.Hour =", time.Hour)
}
