package domain

import (
	"bank/errs"
	"bank/logger"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type customerRepositoryDb struct {
	dbClient *sqlx.DB
}

func NewCustomerRepositoryDb() CustomerRepository {
	dataSourceName := os.Getenv("DB_DATASOURCE")
	if dataSourceName == "" {
		log.Fatal("DB_DATASOURCE environment variable is not set")
	}

	dbClient, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return customerRepositoryDb{dbClient: dbClient}
}

func (d customerRepositoryDb) FindAll(status string) ([]Customer, error) {
	var err error
	customers := []Customer{}

	if status == "" {
		query := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
		err = d.dbClient.Select(&customers, query)
	} else {
		if status == "active" {
			status = "1"
		} else {
			status = "0"
		}
		query := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?"
		err = d.dbClient.Select(&customers, query, status)
	}
	//rows, err := d.dbClient.Query("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?", status)

	/*if err != nil {
		logger.Error("Error querying customers", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	defer rows.Close()

	err = sqlx.StructScan(rows, &customers)*/
	if err != nil {
		logger.Error("Error scanning customers", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	/*for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)
		if err != nil {
			logger.Error("Error scanning customer", zap.Error(err))
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, customer)
	}*/

	return customers, nil
}

func (d customerRepositoryDb) ById(id string) (*Customer, error) {
	customer := &Customer{}
	query := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
	err := d.dbClient.Get(customer, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Customer not found", zap.Error(err))
			return nil, errs.NewNotFoundError("customer not found")
		}
		logger.Error("Error scanning customer", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return customer, nil
}
