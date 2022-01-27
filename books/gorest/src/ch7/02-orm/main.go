package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"carrenolg.io/books/gorest/src/ch7/02-orm/jsonstore/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// DB stores the database session information. Needs to be initialized once.
type DBClient struct {
	db *gorm.DB
}

// UserResponse is the response to be send back for User
type UserResponse struct {
	User models.User `json:"user"`
	Data interface{} `json:"data"`
}

// GetUsersByFirstName fetches the original URL for the give encoded(short) string
func (driver *DBClient) GetUsersByFirstName(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	name := r.FormValue("frist_name")
	// Handle response details
	var query = "SELECT * FROM \"user\" WHERE data->>'frist_name'=?"
	driver.db.Raw(query, name).Scan(&users)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// responseMap := map[string]interface{}{"url": ""}
	respJSON, _ := json.Marshal(users)
	w.Write(respJSON)
}

// GetUser fetches the original URL for the given encoded(short) string
func (driver *DBClient) GetUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	vars := mux.Vars(r)
	// Handle response details
	driver.db.First(&user, vars["id"])
	var userData interface{}
	// Unmarshal JSON string to interface
	json.Unmarshal([]byte(user.Data), &userData)
	var response = UserResponse{User: user, Data: userData}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// responseMap := map[string]interface{}{"url": ""}
	respJSON, _ := json.Marshal(response)
	w.Write(respJSON)
}

// PostUser adds URL to DB and gives back shortened string
func (driver *DBClient) PostUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	postBody, _ := ioutil.ReadAll(r.Body)
	user.Data = string(postBody)
	driver.db.Save(&user)
	responseMap := map[string]interface{}{"id": user.ID}
	var err string = ""
	if err != "" {
		w.Write([]byte("yes"))
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
	dbDriver, _ := dbclient.db.DB()
	defer dbDriver.Close()
	// Create a new router
	r := mux.NewRouter()
	// Attach an elegant path with handler.
	r.HandleFunc("/v1/user/{id:[a-zA-Z0-9]*}", dbclient.GetUser).Methods("GET")
	r.HandleFunc("/v1/user", dbclient.PostUser).Methods("POST")
	r.HandleFunc("/v1/user", dbclient.GetUsersByFirstName).Methods("GET")
	// Serve
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
