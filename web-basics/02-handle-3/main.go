package main

import (
	"fmt"
	"net/http"
)

/*
1. http.HandleFunc
    func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
    type HandlerFunc func(ResponseWriter, *Request)
    func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
    1-1. HandlerFunc可以将某个具有适当签名的函数 f，适配成为一个 Handler，而这个 Handler 具有方法f
   。
*/

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		fmt.Println("Write Error Occurred,", err)
	}
}

func main() {

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: nil,
	}

	http.HandleFunc("/hello", hello)
	//	http.Handle("/hello", http.HandlerFunc(hello))

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Listen Error Occurred,", err)
	}
}
