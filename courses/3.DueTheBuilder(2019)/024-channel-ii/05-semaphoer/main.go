package main

import (
	"fmt"
	"time"
)

func main() {
	// semphoer
	sema := make(chan struct{}, 2)
	canal := make(chan int)

	for i := 0; i < 5; i++ {
		go func(i int) {
			sema <- struct{}{} // get token
			time.Sleep(1 * time.Second)
			canal <- i
			<-sema // return token
		}(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", <-canal)
	}
}
