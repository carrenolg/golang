package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ArticleHandler is a function handler
/*func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}*/

func StrictSlash(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "URL requested by client: %v\n", r.URL.Path)
	path := "/articles/"
	fmt.Fprintf(w, "Server URL : %v\n", path)
}

func main() {
	// Strict Slash
	r := mux.NewRouter()
	r.StrictSlash(true)
	//r.Path("/articles/{category}/{id:[0-9]+}").HandlerFunc(ArticleHandler)
	r.Path("/articles/").HandlerFunc(StrictSlash)

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
