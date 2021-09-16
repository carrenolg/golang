package bank3

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock() // acquire token
	//balance = balance + amount
	defer mu.Unlock() // release token
	deposit(amount)
}

func Balance() int {
	mu.Lock() // acquire token
	b := balance
	mu.Unlock() // release token
	return b
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func deposit(amount int) {
	balance += amount
}
