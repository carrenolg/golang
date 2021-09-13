package main

import "fmt"

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("running foo..")
	}
}

func main() {
	// goroutine
	x := 10

	go foo()

	fmt.Println("Ya terminó foo()", x)
}
