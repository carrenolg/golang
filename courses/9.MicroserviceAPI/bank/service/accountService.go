package service

import (
	"bank/domain"
	"bank/dto"
	"bank/errs"
	"time"
)

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, error)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, error) {

	if err := request.Validate(); err != nil {
		return nil, errs.NewValidationError(err.Error())
	}
	a := domain.Account{
		AccountId:   "",
		CustomerId:  request.CustomerId,
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	account, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	return account.ToDto(), nil

}

func NewAccountService(repo domain.AccountRepository) AccountService {
	return DefaultAccountService{repo: repo}
}
