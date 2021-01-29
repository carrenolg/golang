package main

import "fmt"

func main() {
	// channels
	a := make(chan int, 5)
	b := make(chan int, 5)
	l := make(chan int, 5)

	// all the channels hold so many values until it blocks.
	go func() {
		// hold values
		// a
		a <- 20
		a <- 21
		a <- 22
		// b
		b <- 30
		b <- 31
		b <- 32

		// l
		l <- 1000
		l <- 1001
		l <- 1002
		l <- 1003
		l <- 1004
	}()

	fmt.Println(<-a)
	fmt.Println(<-a)
	fmt.Println(<-a)

	fmt.Println(<-b)
	fmt.Println(<-b)
	fmt.Println(<-b)

	fmt.Println(<-l)
	fmt.Println(<-l)
	fmt.Println(<-l)
	fmt.Println(<-l)
	fmt.Println(<-l)

	//fmt.Println(z1, z2, z3)
}
