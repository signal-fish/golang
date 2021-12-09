package main

import (
	"fmt"
	"net/http"
)

/*
1. http.ListenAndServe()
    1-1. The first parameter is the network address.
        If it is "", then it will listen to port 80 of all network interfaces
    1-2. The second parameter is handler.
        If the value is nil, then it is DefaultServeMux

2. http.Server
    2-1. Addr
        The Addr field indicates the network address.
        If the value is "", it means listen to port 80 of all network interfaces.
    2-2. Handler
        If the Handler field is nil, it means use DefaultServeMux
*/

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			fmt.Println("Write error occurred", err)
		}
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Listen Error Occurred,", err)
	}
}
