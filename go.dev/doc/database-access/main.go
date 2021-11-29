package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "dev",
		Passwd:               "admin",
		Net:                  "tcp",
		Addr:                 "172.19.0.2:3306",
		DBName:               "recordings",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
