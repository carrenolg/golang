package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// concurrency
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Gorutines\t", runtime.NumGoroutine())

	// Increment the WaitGroup counter.
	//wg.Add(1)

	bar()

	wg.Add(1)
	go totausan()
	go foo()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Gorutines\t", runtime.NumGoroutine())

	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	// Decrement the counter when the goroutine completes.
	//defer wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func totausan() {
	for i := 0; i < 100; i++ {
		fmt.Println("totausan:", i)
	}
	defer wg.Done()
}
