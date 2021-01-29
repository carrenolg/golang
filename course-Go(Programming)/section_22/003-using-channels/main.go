package main

import "fmt"

func main() {
	// create a general channel
	ch := make(chan int)

	go send(ch)

	receive(ch)

	fmt.Println("about to exit")

}

// send channel
// you can only push values to the channel
func send(c chan<- int) { c <- 42 }

// receive channel
// you can only pull values from the channel
func receive(c <-chan int) {
	fmt.Println("the value received from channel is:", <-c)
}
