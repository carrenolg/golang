package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"carrenolg.io/books/gorest/src/ch7/01-pq/base62"
	"carrenolg.io/books/gorest/src/ch7/01-pq/models"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

// GenerateShortURL adds URL to DB and gives back shortened string.
func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &record)
	sqlStm := "INSERT INTO web_url(url) VALUES($1) RETURNING id"
	err := driver.db.QueryRow(sqlStm, record.URL).Scan(&id)
	responseMap := map[string]interface{}{"encoded_string": base62.ToBase62(id)}
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBClient{db: db}
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Create a new router
	r := mux.NewRouter()
	// Attach an elegant path with handler.
	r.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]*}", dbclient.GetOriginalURL).Methods("GET")
	r.HandleFunc("/v1/short", dbclient.GenerateShortURL).Methods("POST")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
