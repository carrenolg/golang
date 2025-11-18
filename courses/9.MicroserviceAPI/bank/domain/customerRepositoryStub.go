package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll(status string) ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	return CustomerRepositoryStub{
		customers: []Customer{
			{Id: "1", Name: "John Doe", City: "New York", Zipcode: "10001", DateOfBirth: "1990-01-01", Status: "active"},
			{Id: "2", Name: "Jane Doe", City: "Los Angeles", Zipcode: "90001", DateOfBirth: "1990-01-01", Status: "active"},
		},
	}
}
