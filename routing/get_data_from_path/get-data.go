package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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

	if strings.HasPrefix(request.URL.Path, "/hello/") {
		name := strings.Split(request.URL.Path, "/")[2] // not a good way
		writer.Write([]byte(fmt.Sprintf("Hello, %s", name)))
		return
	}
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("404 Page Not Found"))
}

func main() {
	err := http.ListenAndServe("localhost:3000", MyHandler{})
	check(err)
}
