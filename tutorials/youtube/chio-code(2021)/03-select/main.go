package main

import (
	"fmt"
	"time"
)

func main() {
	// select
	c1 := make(chan string)
	c2 := make(chan string)
	quit := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one" // send
		close(c1)
		quit <- true
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two" //sends
		close(c2)
	}()

	// receive
	for {
		select {
		case msg1, ok := <-c1:
			if ok {
				fmt.Println("received", msg1)
			}
		case msg2, ok := <-c2:
			if ok {
				fmt.Println("received", msg2)
			}
		case <-quit:
			fmt.Println("quit")
			return // break loop
		}

	}
}
