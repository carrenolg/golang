package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("running..")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
		// c2 <- "two"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		// c1 <- "one"
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
