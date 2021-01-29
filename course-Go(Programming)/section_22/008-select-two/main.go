// Go program to illustrate the
// concept of select statement
package main

import (
	"fmt"
	"time"
)

// main function
func main() {

	// Creating channels
	ch1 := make(chan string)
	ch2 := make(chan string)
	quit := make(chan int)

	// calling function 1 and
	// function 2 in goroutine
	go portal1(ch1, quit)
	go portal2(ch2, quit)

	receive(ch1, ch2, quit)

}

// function 1
func portal1(ch1 chan<- string, q chan<- int) {
	time.Sleep(1 * time.Second)
	ch1 <- "Welcome to channel 1"
	//close(ch1)
}

// function 2
func portal2(ch2 chan<- string, q chan<- int) {
	time.Sleep(3 * time.Second)
	ch2 <- "Welcome to channel 2"
	//close(ch2)
	q <- 0
}

func receive(ch1, ch2 <-chan string, q <-chan int) {
	// select is waiting that the communication starts
	for {
		select {
		// case 1 for portal 1
		case op1 := <-ch1:
			fmt.Println(op1)
		// case 2 for portal 2
		case op2 := <-ch2:
			fmt.Println(op2)
		case <-q:
			return
		}
	}

}
