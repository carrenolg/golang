package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	// Fecth query parms as a map
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category:%s\n", queryParams["category"][0])
}

func main() {
	// create a new router
	r := mux.NewRouter()
	// attach a path with handler
	r.HandleFunc("/articles", QueryHandler)
	r.Queries("id", "category")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	// http://localhost:8000/articles?id=1345&category=birds
}
