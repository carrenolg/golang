package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan struct{})

	go func() {
		// wait
		time.Sleep(2 * time.Second)
		ch <- struct{}{}
	}()
	fmt.Println("waiting...")
	<-ch
	fmt.Println("exit program")
}
