package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST: Customer created")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Customer ID: %s\n", vars["customer_id"])
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John", City: "New York", Zipcode: "10001"},
		{Name: "Jane", City: "Los Angeles", Zipcode: "90001"},
		{Name: "Jim", City: "Chicago", Zipcode: "60601"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
