package main

import "fmt"

func main() {
	// channel
	ch := make(chan string)
	// send
	go func() {
		ch <- "meli"
		ch <- "meli2"
	}()
	// receive
	output := <-ch
	output2 := <-ch
	fmt.Println(output)
	fmt.Println(output2)
}
