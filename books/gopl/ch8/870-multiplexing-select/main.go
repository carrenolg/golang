package main

import (
	"fmt"
)

func main() {
	// channel with buffer
	ch := make(chan int, 5)

	for i := 0; i < 10; i++ {
		// fmt.Println("loop", i)
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			// do nothing
		}
	}
}
