package main

import (
	"fmt"
)

func main() {
	// send
	c := gen()
	// receive
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	// you need use a goroutine to send values into channel
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	// pull the values off
	for v := range c {
		fmt.Println(v)
	}
}
