package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(rw, "Method not allowed")
		return
	}

	fmt.Fprintf(rw, "Hello there %s", "visitor\n")
}

func getCountries(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(countries)

}

func addCountry(wr http.ResponseWriter, r *http.Request) {

	country := &Country{}

	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(wr, "%v", err)
		return
	}

	countries = append(countries, country)
	fmt.Fprintf(wr, "country was added")

}
