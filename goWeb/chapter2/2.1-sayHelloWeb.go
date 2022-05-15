package main

import (
	"net/http"
)

func sayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":8082", nil)
}