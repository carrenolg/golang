package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
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

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("Intro text not correct; got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	input := strings.NewReader("q")
	reader := bufio.NewScanner(input)
	res, _ := checkNumbers(reader)

	if !strings.EqualFold(res, "") {
		t.Error("Keyword worng for exit program!")
	}

	// convert user input to int
	input = strings.NewReader("1.0")
	reader = bufio.NewScanner(input)
	res, _ = checkNumbers(reader)
	if res != "Please enter a whole number!" {
		t.Error("Input can't be converted into integer")
	}

	input = strings.NewReader("7")
	reader = bufio.NewScanner(input)
	res, expected := checkNumbers(reader)
	if expected {
		t.Errorf("expected '7 is a prime number!', and got %s", res)
	}
}

func Test_ReadUserInput(t *testing.T) {
	// create chan
	doneChan := make(chan bool)

	// create buffer
	var stdin bytes.Buffer

	// write in buffer
	stdin.Write([]byte("1\nq\n"))

	// launch goroutine
	go readUserInput(&stdin, doneChan)

	// wait to goroutine ending
	<-doneChan
	close(doneChan)
}
