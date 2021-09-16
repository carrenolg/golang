package main

import (
	"fmt"

	bank "carrenolg.io/books/gopl/ch9/bank1"
)

type Cake struct {
	state string
}

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // baker never touches this cake again
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer never touches this cake again
	}
}

func main() {
	// 9.1 Race Conditions

	// second way to avoid a data race (monitor goroutine)
	bank.Deposit(100)
	fmt.Println(bank.Balance())

	// serial confinement
	cooked := make(chan *Cake)
	iced := make(chan *Cake)

	go baker(cooked)
	go icer(iced, cooked)

	for cake := range iced {
		fmt.Println((*cake).state) // infinite loop
	}

}
