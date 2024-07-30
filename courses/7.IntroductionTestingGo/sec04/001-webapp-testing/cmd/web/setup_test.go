package main

import (
	"log"
	"os"
	"testing"
	"webapp/pkg/db"
)

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	// set up session
	app.Session = getSession()
	app.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5"

	// db
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DBConn = db.PostgresConn{DB: conn}

	os.Exit(m.Run())
}
