package domain

import (
	"bank/dto"
	"encoding/json"
	"encoding/xml"
)

type Customer struct {
	Id          string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"name" xml:"name" db:"name"`
	City        string `json:"city" xml:"city" db:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode" db:"zipcode"`
	DateOfBirth string `json:"date_of_birth" xml:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status" db:"status"`
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, error)
	ById(id string) (*Customer, error)
}

func (c Customer) ToJson() string {
	json, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(json)
}

func (c Customer) ToXml() string {
	xml, err := xml.Marshal(c)
	if err != nil {
		return ""
	}
	return string(xml)
}

func (c Customer) ToDto() *dto.CustomerResponse {

	return &dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.StatusAsText(),
	}
}

func (c Customer) StatusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}
