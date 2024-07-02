package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	DSN     string
	DB      *sql.DB
	Session *scs.SessionManager
}

func main() {
	// set up an application config
	app := application{}

	// set up a database connection
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection")
	flag.Parse()

	// connect to the database
	db, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app.DB = db

	// get a session manager
	app.Session = getSession()

	// print out a message server starting
	// UTC time, server default
	log.Println("Starting server on port 8080...")

	// start the server
	err = http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal()
	}
}
