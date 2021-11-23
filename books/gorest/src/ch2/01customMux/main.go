package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// CustomServer is a struct which can be a multiplexer
type CustomServerMux struct {
}

// This is the function handler to be overridden
func (p *CustomServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f\n", rand.Float64())
}

func main() {
	// Any struct that has serveHTTP function can be a multiplexer
	mux := &CustomServerMux{}
	http.ListenAndServe(":8000", mux)
}
