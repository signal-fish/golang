package main

import (
	"flag"
	"fmt"
)

func main() {
	var user, pwd, host string
	var port int

	flag.StringVar(&user, "u", "", "User name, default empty")
	flag.StringVar(&pwd, "pwd", "", "Password, default empty")
	flag.StringVar(&host, "h", "localhost", "Host")
	flag.IntVar(&port, "p", 3306, "Port")

	flag.Parse()
	fmt.Printf("user = %v\npwd = %v\nhost = %v\nport = %v\n", user, pwd, host, port)
}
