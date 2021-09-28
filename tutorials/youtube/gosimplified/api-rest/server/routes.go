package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)

	http.HandleFunc("/countries", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCountries(rw, r)

		case http.MethodPost:
			addCountry(rw, r)

		default:
			if r.Method != http.MethodGet {
				rw.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(rw, "Method not allowed")
				return
			}
		}
	})
}
