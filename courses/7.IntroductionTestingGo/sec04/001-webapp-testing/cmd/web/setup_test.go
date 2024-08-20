package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	// set up session
	app.Session = getSession()

	app.DBConn = &dbrepo.TestDBRepo{}

	os.Exit(m.Run())
}
