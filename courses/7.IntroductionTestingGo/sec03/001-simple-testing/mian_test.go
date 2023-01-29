package main

import (
	"io"
	"os"
	"testing"
)

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false.", 0)
	}
	if msg != "0 is not prime, by definition!" {
		t.Error("Wrong message returned:", msg)
	}
	result, msg = isPrime(7)
	if !result {
		t.Errorf("with %d as test parameter, got false, but expected true.", 7)
	}
	if msg != "7 is a prime number!" {
		t.Error("Wrong message returned:", msg)
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}

}
