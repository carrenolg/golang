package domain

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type customerRepositoryDb struct {
	dbClient *sql.DB
}

func (d customerRepositoryDb) FindAll(status string) ([]Customer, error) {
	customers := []Customer{}
	if status == "active" {
		status = "1"
	} else {
		status = "0"
	}
	rows, err := d.dbClient.Query("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?", status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepository {
	dataSourceName := os.Getenv("DB_DATASOURCE")
	if dataSourceName == "" {
		log.Fatal("DB_DATASOURCE environment variable is not set")
	}

	dbClient, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return customerRepositoryDb{dbClient: dbClient}
}
