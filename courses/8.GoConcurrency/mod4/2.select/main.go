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
		for i := range 10 {
			time.Sleep(1 * time.Second)
			c1 <- fmt.Sprintf("[%d]: Go1", i)
		}
		close(c1)
	}()

	go func() {
		for i := range 2 {
			time.Sleep(3 * time.Second)
			c1 <- fmt.Sprintf("[%d]: Go2", i)
		}
	}()

	for {
		select {
		case msg1, ok := <-c1:
			if !ok {
				fmt.Println("c2 channel closed")
				return
			}
			fmt.Println("received", msg1)
		case msg2, ok := <-c2:
			if !ok {
				fmt.Println("c2 channel closed")
				return
			}
			fmt.Println("received", msg2)
		}
	}
}
