package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
	msg string
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		// re-check avoid race condition
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {
	newSingle := getInstance()
	newSingle.msg = "one instance"
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
