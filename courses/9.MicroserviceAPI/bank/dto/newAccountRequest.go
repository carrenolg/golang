package dto

import "bank/errs"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type NewAccountResponse struct {
	AccountId string `json:"account_id"`
}

func (r NewAccountRequest) Validate() error {
	if r.Amount < 5000 {
		return errs.NewValidationError("amount must be greater than 5000")
	}
	if r.AccountType != "saving" && r.AccountType != "checking" {
		return errs.NewValidationError("account_type must be either savings or checking")
	}
	return nil
}
