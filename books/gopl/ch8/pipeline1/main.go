package main

import "fmt"

func main() {
	// pipeline1
	naturals := make(chan int)
	squares := make(chan int)

	// Counters
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}

}
