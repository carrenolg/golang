package main

import "fmt"

func main() {
	// channel
	ch := make(chan string)
	// send
	go func() {
		ch <- "hello"
		ch <- "how are you?"
	}()
	// received
	received := <-ch // bloking until send data
	received = <-ch  // bloking until send data
	fmt.Println("I'd received:", received)
}
