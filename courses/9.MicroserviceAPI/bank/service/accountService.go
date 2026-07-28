package service

import (
	"bank/domain"
	"bank/dto"
	"bank/errs"
	"time"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, error)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, error)
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

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, error) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindById(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
