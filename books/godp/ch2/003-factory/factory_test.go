package creational

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' musg be exist.")
	}
	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct.")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCar(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method of type 'DebitCar' musg be exist.")
	}
	msg := payment.Pay(22.30)

	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct.")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodCreditCar(t *testing.T) {
	payment, err := GetPaymentMethod(CreditCard)
	if err != nil {
		t.Fatal("A payment method of type 'CreditCar' must be exist.")
	}
	msg := payment.Pay(22.30)

	if !strings.Contains(msg, "paid using credit card") {
		t.Error("The credit card payment method message wasn't correct.")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExist(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Fatal("A payment method with ID 20 must return an error.")
	}
	t.Log("LOG:", err)
}
