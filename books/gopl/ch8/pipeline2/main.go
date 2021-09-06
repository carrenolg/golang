package main

import "fmt"

func main() {
	// pipeline2
	naturals := make(chan int)
	squares := make(chan int)

	// Counters
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals) // all data have been sent
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares) // all data have been sent
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}

}
