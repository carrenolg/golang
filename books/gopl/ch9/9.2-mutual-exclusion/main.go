package main

import (
	"fmt"

	"carrenolg.io/books/gopl/ch9/bank2"
	"carrenolg.io/books/gopl/ch9/bank3"
)

func main() {
	// 9.2 Mutual Exclusion

	// using semphore
	bank2.Deposit(500)
	bank2.Deposit(300)
	bank2.Deposit(-400)
	fmt.Println(bank2.Balance())

	// using asyn.Mutex
	bank3.Deposit(500)
	bank3.Deposit(-500)
	fmt.Println(bank3.Balance())

	// using asyn.Mutex (complex)
	// atomic operation (no re-entrant)
	trans := bank3.Withdraw(100)
	fmt.Println(trans)

}
