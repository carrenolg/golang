package service

import (
	"bank/domain"
	"bank/dto"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, error)
	GetCustomer(id string) (*dto.CustomerResponse, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, error) {
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	customerResponses := make([]dto.CustomerResponse, len(customers))
	for i, customer := range customers {
		customerResponses[i] = *customer.ToDto()
	}
	return customerResponses, nil
}

func NewCustomerService(repo domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repo: repo}
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, error) {
	customer, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	return customer.ToDto(), nil
}
