package domain

import (
	"bank/errs"
	"bank/logger"
	"strconv"
	"time"

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

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, error) {
	// starting the database transaction

	tx, beginErr := d.client.Begin()
	if beginErr != nil {
		logger.Error("Error starting database transaction", zap.Error(beginErr))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// inserting bank transaction
	result, insertErr := tx.Exec("INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)", t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating bank account balance
	if t.IsWithdrawal() {
		_, insertErr = tx.Exec("UPDATE accounts SET amount = amount - ? WHERE account_id = ?", t.Amount, t.AccountId)
	} else {
		_, insertErr = tx.Exec("UPDATE accounts SET amount = amount + ? WHERE account_id = ?", t.Amount, t.AccountId)
	}

	// in case of error, rollback the transaction
	if insertErr != nil {
		tx.Rollback()
		logger.Error("Error updating bank account balance", zap.Error(insertErr))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// getting the last insert id
	lastInsertId, lastInsertErr := result.LastInsertId()
	if lastInsertErr != nil {
		tx.Rollback()
		logger.Error("Error getting last insert id", zap.Error(lastInsertErr))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	t.TransactionId = strconv.FormatInt(lastInsertId, 10)

	// committing the transaction before reading updated balance
	if commitErr := tx.Commit(); commitErr != nil {
		logger.Error("Error committing database transaction", zap.Error(commitErr))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// getting the updated bank account balance after commit
	account, accountErr := d.FindById(t.AccountId)
	if accountErr != nil {
		logger.Error("Error getting bank account", zap.Error(accountErr))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	t.Amount = account.Amount
	t.TransactionDate = time.Now().Format("2006-01-02 15:04:05")
	return &t, nil

}

func (d AccountRepositoryDb) FindById(accountId string) (*Account, error) {
	query := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts WHERE account_id = ?"
	var account Account
	err := d.client.Get(&account, query, accountId)
	if err != nil {
		logger.Error("Error getting bank account", zap.Error(err))
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepository {
	return AccountRepositoryDb{client: dbClient}
}
