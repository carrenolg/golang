package main

import "testing"

func Test_table_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"prime", 8, false, "8 is not a prime number because it is divisible by 2"},
		{"prime", 13, true, "13 is a prime number!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative", -11, false, "Negative numbers are not prime, by definition!"},
	}
	// tests
	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected != result {
			t.Errorf("with %d as test parameter, got %v, but expected %v", e.testNum, result, e.expected)
		}
		if msg != e.msg {
			t.Errorf("%s: expected %s but got %s.", e.name, e.msg, msg)
		}
	}
}
