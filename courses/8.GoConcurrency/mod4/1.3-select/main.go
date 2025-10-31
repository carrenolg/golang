package main

import (
	"fmt"
	"time"
)

func main() {
	// mod4 1.1 - select in/out channels

	inchan := make(chan string)
	outchan := make(chan string)

	go func() {
		for i := range 10 {
			time.Sleep(1 * time.Second)
			inchan <- fmt.Sprintf("[%d]: Go1", i)
		}
		close(inchan)
	}()

	go func() {
		for msg := range outchan {
			fmt.Printf("%s\n", msg)
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case msg1, ok := <-inchan:
			if !ok {
				fmt.Println("c1 channel closed")
				return
			}
			fmt.Println("received", msg1)
		case outchan <- "ping":
		}
	}
}
