package main

import "fmt"

func main() {
	// channel
	ch := make(chan int, 3)
	// send
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	// iterate
	for num := range ch {
		fmt.Println("Receiving:", num)
	}
	fmt.Println("Channel closed.")
}
