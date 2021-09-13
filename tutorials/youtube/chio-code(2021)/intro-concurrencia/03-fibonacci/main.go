package main

import (
	"fmt"
)

func fibo(numbers, done chan int) {
	x, y := 0, 1
	for {
		select {
		case numbers <- x:
			x, y = y, x+y
		case <-done:
			fmt.Println("done!")
			return // signal goroutine to exit
		}
	}
}

func main() {
	// select statement
	numbers := make(chan int)
	done := make(chan int)

	// call fibo
	go fibo(numbers, done)

	// receive
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-numbers)
		}
		done <- 0 // cancel signal (goroutine exit)
	}()
	// fmt.Println(<-numbers) // error
}
