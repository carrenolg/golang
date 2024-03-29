package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Book is a placeholder for book
type Book struct {
	Id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	log.Println(db)
	if err != nil {
		log.Println(err)
	}
	// Create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table books.")
	}
	statement.Exec()
	// Create
	statement, _ = db.Prepare("INSERT INTO books (name, author, isbn) VALUES(?, ?, ?)")
	statement.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	log.Println("Inserted the book into database")

	// Read
	rows, _ := db.Query("SELECT id, name, author FROM books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.Id, &tempBook.name, &tempBook.author)
		// log
		log.Printf("ID:%d, Book:%s, Author:%s", tempBook.Id, tempBook.name, tempBook.author)
	}

	// Update
	statement, _ = db.Prepare("UPDATE books SET name=? WHERE id=?")
	statement.Exec("The Tale of Two Cities", 1)
	log.Println("Sucessfully updated the book in database.")

	// Delete
	statement, _ = db.Prepare("DELETE FROM books WHERE id=?")
	statement.Exec(3)
	log.Println("Sucessfully deleted the book in database.")

}
