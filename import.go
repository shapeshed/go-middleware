package main

import (
	"./middleware"
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	finalHandler := http.HandlerFunc(HelloServer)
	http.Handle("/hello", middleware.SetContentType(middleware.SetApiVersion(finalHandler)))

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
