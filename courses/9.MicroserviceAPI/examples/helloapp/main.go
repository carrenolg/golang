package main

import (
	"fmt"
	"net/http"
)

func main() {

	// define routes
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	// start server
	http.ListenAndServe(":8080", nil)
}
