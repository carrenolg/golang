package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"
)

type City struct {
	Name string
	Area uint64
}

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		// Filtering requests by MIME type
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Please send JSON."))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		// Setting cookie to each and every response
		cookie := http.Cookie{
			Name:  "Server-Time(UTC)",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
		}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Check if method is POST
	if r.Method == "POST" {
		var tempCity City
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		// Create response
		fmt.Printf("Got %s city with area of %d sq miles\n", tempCity.Name, tempCity.Area)
		// Tell everything is fine
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		// Say method not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	chain := alice.New(filterContentType, setServerTimeCookie).Then(mainLogicHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8000", nil)
}
