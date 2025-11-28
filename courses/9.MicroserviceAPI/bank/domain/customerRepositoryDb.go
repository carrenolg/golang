package domain

import (
	"bank/errs"
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
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	defer rows.Close()

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)
		if err != nil {
			return nil, errs.NewUnexpectedError("unexpected database error")
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

func (d customerRepositoryDb) ById(id string) (*Customer, error) {
	customer := &Customer{}
	row := d.dbClient.QueryRow("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?", id)
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return customer, nil
}
