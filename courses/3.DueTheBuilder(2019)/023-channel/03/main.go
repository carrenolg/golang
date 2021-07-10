package main

import (
	"fmt"
	"time"
)

func double(n int) int {
	time.Sleep(500 * time.Millisecond)
	return n * 2
}

func triple(n int) int {
	time.Sleep(800 * time.Millisecond)
	return n * 3
}

func main() {
	// channels
	dobles := make(chan int)
	triples := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			dobles <- double(i)
			triples <- triple(i)
		}
		close(dobles)
		close(triples)
	}()

loop:
	for {
		select {
		case i, ok := <-dobles:
			if !ok {
				break loop
			}
			fmt.Println(i, " ")
		case i, ok := <-triples:
			if !ok {
				break loop
			}
			fmt.Println(i, " ")
		}
	}
}
