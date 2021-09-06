package main

import "fmt"

func main() {
	// 8.4.4 Buffered Channels
	ch := make(chan string, 3)

	ch <- "A"
	ch <- "B"
	ch <- "C"

	// receive one value
	fmt.Println(<-ch) // "A"

	// buffer capacity
	fmt.Println(cap(ch)) // "3"

	// elements into the buffer
	fmt.Println(len(ch)) // "2"

	fmt.Println(<-ch) // "B"
	fmt.Println(<-ch) // "C"
}
