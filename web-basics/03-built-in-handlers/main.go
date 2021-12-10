package main

import (
	"fmt"
	"net/http"
)

func main() {
	path := "/home/signalfish/golang-basics/web-basics/03-built-in-handlers"
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, path+r.URL.Path)
		})
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			fmt.Println("Listen Error Occurred:", err)
		}
	*/
	err := http.ListenAndServe(":8000", http.FileServer(http.Dir(path)))
	if err != nil {
		fmt.Println("Listen Error Occurred:", err)
	}
}
