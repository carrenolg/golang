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
	// mux.Vars
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Setting ID is: %v\n", vars["id"])
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Details ID is: %v\n", vars["id"])
}

func main() {
	// Define path
	r := mux.NewRouter()
	r.Path("/articles/{category}/{id:[0-9]+}").HandlerFunc(ArticleHandler)
	// Subpath
	s := r.PathPrefix("/articles").Subrouter()
	// Handlers
	s.HandleFunc("/", ArticleHandler)               // "/articles/"
	s.HandleFunc("/{id}/settings", SettingsHandler) // "/articles/123/settings"
	s.HandleFunc("/{id}/details", DetailsHandler)   // "/articles/123/details"

	// Static path
	path := "/home/gio10/cs/dev/golang/books/gorest/src/ch2/06-customPath/tmp/static"
	static := http.StripPrefix("/static/", http.FileServer(http.Dir(path)))
	r.PathPrefix("/static/").Handler(static)

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
