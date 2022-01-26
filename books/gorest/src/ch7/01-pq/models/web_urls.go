package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	var err error
	strConn := "postgres://admin:1q2w3e4r@172.19.0.2/testdb?sslmode=disable"
	db, err := sql.Open("postgres", strConn)
	if err != nil {
		return nil, err
	} else {
		// Create model for our URL service
		/*sqlCreate := `CREATE TABLE WEB_URL(
			ID SERIAL PRIMARY KEY,
			URL TEXT NOT NULL
		);`
		stmt, err := db.Prepare(sqlCreate)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		res, err := stmt.Exec()
		log.Println(res)
		if err != nil {
			log.Println(err)
			return nil, err
		}*/
		return db, nil
	}
}
