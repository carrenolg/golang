package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers", getAllCustomers)
	// start server
	http.ListenAndServe(":8080", router)
}
