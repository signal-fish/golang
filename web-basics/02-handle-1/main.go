package main

import (
	"fmt"
	"net/http"
)

/*
1. http.Server
    1-1. Addr
        The Addr field indicates the network address.
        If the value is "", it means listen to port 80 of all network interfaces.
    1-2. Handler
        If the Handler field is nil, it means use DefaultServeMux
2. Handler
    type Handler interface {
	ServeHTTP(ResponseWriter,  *Request)
    }
    2-1. Handler is an interface
    2-2. Handler defines a method ServeHTTP() with two parameters, one is HTTPResponseWriter, the other is a
pointer to the Request struct.
*/

type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println("Error occurred,", err)
	}
}

func main() {
	mh := myHandler{}
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: &mh,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Listen error occurred,", err)
	}
}
