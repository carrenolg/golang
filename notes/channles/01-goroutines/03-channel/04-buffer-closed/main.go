package main

import "fmt"

func main() {
	// buffer
	ch := make(chan int, 3)
	// send
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	// receive
	for i := 0; i < 5; i++ {
		fmt.Println("receiving:", <-ch)
	}
}
