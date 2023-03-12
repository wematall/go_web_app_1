package main

import (
	"log"
	"net/http"
)

type MyHandler struct {
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		writer.Write([]byte("Home"))
		return
	}

	if request.URL.Path == "/hello" {
		writer.Write([]byte("Hello page"))
		return
	}

	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("404 Page Not Found"))
}

func main() {
	err := http.ListenAndServe("localhost:3000", MyHandler{})
	check(err)
}
