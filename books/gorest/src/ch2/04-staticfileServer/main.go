package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("/home/gio10/cs/dev/golang/books/gorest/src/ch2/04-staticfileServer/static"))

	log.Fatal(http.ListenAndServe(":8000", router))
}
