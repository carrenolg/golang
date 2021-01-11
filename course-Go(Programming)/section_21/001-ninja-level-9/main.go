package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())
	// exercise #1

	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("cpu", runtime.NumCPU())
	fmt.Println("gs", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("cpu", runtime.NumCPU())
	fmt.Println("gs", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	defer wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
	defer wg.Done()
}
