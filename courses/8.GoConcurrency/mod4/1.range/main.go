package main

import "fmt"

var ch = make(chan int)

func main() {
	fmt.Println("Run range ...")

	// sender
	go sender()

	// receiver
	for i := range ch {
		fmt.Println(i)
	}
}

func sender() {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
