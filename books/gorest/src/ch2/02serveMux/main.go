package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	// create a new ServeMux
	newMux := http.NewServeMux()

	// attach multiple handlers to ServeMux
	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})

	// create server
	http.ListenAndServe(":8000", newMux)
}
