package main

import "fmt"

// numbers
const nums = 3

func main() {
	// channel
	ch := make(chan int)
	// send
	go sender(ch)
	// receive
	receiver(ch)

}

func sender(ch chan<- int) {
	for i := 1; i <= nums; i++ {
		ch <- i
		fmt.Println(i, "sending successful")
	}
}

func receiver(ch <-chan int) {
	for i := 1; i <= nums; i++ {
		num := <-ch
		fmt.Println("received:", num)
	}
}
