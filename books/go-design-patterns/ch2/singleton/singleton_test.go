package main

import "testing"

func TestGetInstance(t *testing.T) {
	count := GetInstance()
	if count == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}
	currentCounter := count.AddOne()
	if currentCounter != 1 {
		t.Errorf("After calling for the first time to count.AddOne(), the count must be 1 but it is %d\n", currentCounter)
	}
	newCounter := count.AddOne()
	if newCounter != 2 {
		t.Errorf("After calling for the second time to count.AddOne(), the count must be 2 but it is %d\n", newCounter)
	}
}
