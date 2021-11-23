package main

import (
	"io"
	"log"
	"net/http"
)

func MyServe(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", MyServe)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
