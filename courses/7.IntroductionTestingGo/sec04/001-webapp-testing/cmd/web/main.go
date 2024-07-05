package main

import (
	"flag"
	"log"
	"net/http"
	"webapp/pkg/db"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	DSN     string
	DBConn  db.PostgresConn
	Session *scs.SessionManager
}

func main() {
	// set up an application config
	app := application{}
	dbPosgresConn := db.PostgresConn{}

	// set up a database connection
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection")
	flag.Parse()

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// set up the database connection
	dbPosgresConn.DB = conn
	app.DBConn = dbPosgresConn

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
