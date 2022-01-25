package main

import (
	"database/sql"
	"net/http"

	"carrenolg.io/books/gorest/src/ch7/01-pq/base62"
	"github.com/gorilla/mux"
)

// DB stores the database session information. Needs to be initializaed
type DBClient struct {
	db *sql.DB
}

// Model the record struct
type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

// GetOriginalURL fetches the original URL for the given encoded (short) string.
func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)
	// Get ID from base62 string
	id := base62.ToBase10(vars["encoded_string"])
	stm := "SELECT url FROM web_url WHERE ID = $1"
	err := driver.db.QueryRow(stm, id).Scan(&url)
	// Handle response details
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

	}

}

func main() {

}
