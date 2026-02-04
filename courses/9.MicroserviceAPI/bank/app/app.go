package app

import (
	"bank/domain"
	"bank/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// get db client
	dbClient := getDbClient()

	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb(dbClient))}
	ah := AccountHandlers{service: service.NewAccountService(domain.NewAccountRepositoryDb(dbClient))}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// account routes
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

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

func getDbClient() *sqlx.DB {
	dataSourceName := os.Getenv("DB_DATASOURCE")
	if dataSourceName == "" {
		log.Fatal("DB_DATASOURCE environment variable is not set")
	}
	dbClient, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)
	dbClient.SetConnMaxLifetime(time.Duration(10) * time.Minute)

	return dbClient
}
