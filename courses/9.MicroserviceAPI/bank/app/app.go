package app

import (
	"bank/domain"
	"bank/service"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// start server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = ":8080"
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)
}
