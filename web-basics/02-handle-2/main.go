package main

import (
	"fmt"
	"net/http"
)

/*
1. http.Handle
    func Handle(pattern string, handler Handler)
    type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
    }
    1-1. http 包有一个 Handle 函数，ServerMux struct 也有一个 Handle 方法。调用 http.Handle，实际上调用的是 DefaultServeMux 上的 Handle 方法。
    1-2. Handler 函数的签名与 ServeHTTP 方法的签名一样，都是接收http.ResponseWriter和*http.Request 两个参数。
*/

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		fmt.Println("Write Error Occurred")
	}
}

type aboutHandler struct{}

func (h *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("About"))
	if err != nil {
		fmt.Println("Write Error Occurred")
	}
}

func main() {
	h := helloHandler{}
	a := aboutHandler{}

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: nil,
	}

	http.Handle("/hello", &h)
	http.Handle("/about", &a)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Listen Error occurred,", err)
	}
}
