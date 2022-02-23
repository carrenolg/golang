package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)

	// g1
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("run g1")
				time.Sleep(time.Millisecond * 1000 / 4)
			}
		}
	}()
	fmt.Println("run g0")
	time.Sleep(time.Second * 1)
	quit <- true
}
