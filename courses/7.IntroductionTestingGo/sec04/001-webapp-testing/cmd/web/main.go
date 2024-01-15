package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	// set up an application config
	app := application{}

	// get a session manager
	app.Session = getSession()

	// print out a message server starting
	// UTC time, server default
	log.Println("Starting server on port 8080...")

	// start the server
	err := http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal()
	}
}
