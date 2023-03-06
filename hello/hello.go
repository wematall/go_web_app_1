package main

import (
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type MyHandler struct {
}

func (MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, World!\n"))
}

func main() {
	err := http.ListenAndServe("localhost:3000", MyHandler{})
	check(err)
}
