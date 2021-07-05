package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// goroutine
	// info
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU's\t", runtime.NumCPU())
	fmt.Println("Goroutine\t", runtime.NumGoroutine())

	// run
	wg.Add(1)
	go foo(1)
	bar()
	// foo and bar print out their data concurrently

	fmt.Println("CPU's\t", runtime.NumCPU())
	fmt.Println("Goroutine\t", runtime.NumGoroutine())

	wg.Wait()
}

func foo(n int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", n, i)
	}
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("bar:", i)
	}
}
