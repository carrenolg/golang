package domain

import "bank/dto"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(account Account) (*Account, error)
	SaveTransaction(transaction Transaction) (*Transaction, error)
	FindById(id string) (*Account, error)
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
