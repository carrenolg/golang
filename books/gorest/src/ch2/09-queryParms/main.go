package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ArticleHandler is a function handler
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// Query parms
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter %v\n", queryParams["category"])
	fmt.Fprintf(w, "Got parameter %v\n", queryParams["id"])
	fmt.Fprintf(w, "Got parameter %v\n", queryParams["area"])
}

func main() {
	// Query parameters
	r := mux.NewRouter()
	r.HandleFunc("/articles", ArticleHandler)
	r.Queries("id", "category")
	// Example:
	// curl -X GET 'http://localhost:8000/articles?id=123&category=books&area=science'

	// Server
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
