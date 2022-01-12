package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	strConnection := "root:1q2w3e4r@tcp(172.17.0.3)/books?parseTime=true"
	db, err := sql.Open("mysql", strConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Test Connection
	db.SetMaxOpenConns(5)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Ping to db sucessful!")
	}

	// Get books by id
	query := `SELECT id, title, isbn, release_date FROM item WHERE id=?`
	var book struct {
		ID          string
		Title       string
		ISBN        string
		ReleaseDate time.Time
	}
	err = db.QueryRow(query, 1).Scan(&book.ID, &book.Title, &book.ISBN, &book.ReleaseDate)
	if err != nil {
		log.Println("query error:", err)
	} else {
		fmt.Println("row", book)
	}

	// Get all books from the item table
	query = `SELECT id, title, isbn, release_date FROM item`
	type Item struct {
		ID          string
		Title       string
		ISBN        string
		ReleaseDate time.Time
	}
	var books []Item
	rows, err := db.Query(query)
	if err != nil {
		log.Println("query error:", err)
	}
	for rows.Next() {
		var b Item
		err := rows.Scan(&b.ID, &b.Title, &b.ISBN, &b.ReleaseDate)
		if err != nil {
			log.Println("scan error:", err)
		} else {
			books = append(books, b)
		}
	}
	fmt.Println("rows", books)

	// Create new book
	query = `INSERT INTO item ( title, isbn, release_date) VALUES(?, ?, ?)`
	stm, err := db.Prepare(query)
	if err != nil {
		log.Println("prepare error", err)
	}
	result, err := stm.Exec("El cid", "ISB462CDF", "1995-06-06")
	if err != nil {
		log.Println(err)
	} else {
		id, _ := result.LastInsertId()
		fmt.Println("book-id:", id)
	}
}
