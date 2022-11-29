package main

import "fmt"

func main() {
	// buffer
	//ch := make(chan string) // blocking on execution
	ch := make(chan string, 1)
	ch <- "hola"
	received := <-ch
	fmt.Println(received)
}
