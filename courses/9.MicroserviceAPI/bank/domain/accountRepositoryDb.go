package domain

import (
	"bank/errs"
	"bank/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(account Account) (*Account, error) {
	query := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(query, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error("Error saving account", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error getting last insert id", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	account.AccountId = strconv.FormatInt(id, 10)
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepository {
	return AccountRepositoryDb{client: dbClient}
}
