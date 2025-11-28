package service

import "bank/domain"

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, error) {
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func NewCustomerService(repo domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repo: repo}
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
	customer, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
