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
	var cancela chan struct{}
	// event
	go func() {
		time.Sleep(5 * time.Second)
		cancela = make(chan struct{})
		cancela <- struct{}{}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			dobles <- double(i)  // send
			triples <- triple(i) // send
		}
		close(dobles)
		close(triples)
	}()

loop:
	for {
		select {
		// receive
		case i, ok := <-dobles:
			if !ok {
				break loop
			}
			fmt.Println(i, " ")
		// receive
		case i, ok := <-triples:
			if !ok {
				break loop
			}
			fmt.Println(i, " ")
		// get notification
		case <-cancela:
			fmt.Println("Cancelado!")
			break loop
		}
	}
}
