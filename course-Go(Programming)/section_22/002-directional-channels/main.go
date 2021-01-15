package main

import "fmt"

func main() {
	// create channel
	c := make(chan int, 2)

	c <- 42
	fmt.Println(<-c)

	c <- 43
	c <- 44
	fmt.Println(<-c)
	fmt.Println(<-c)

	// type
	fmt.Printf("%T\n", c) // [chan int]

	fmt.Println("--------")

	// send and receive (bi-directional)
	ch := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Printf("channel:\t\t%T\n", ch)
	fmt.Printf("channel receive:\t%T\n", cr)
	fmt.Printf("channel send:\t\t%T\n", cs)

	fmt.Println("--------")
	// check channel send
	chSend := make(chan<- int)

	go func() {
		chSend <- 42
	}()

	// invalid operation: <-chSend (receive from send-only type chan<- int)
	//fmt.Println(<-chSend)

	fmt.Printf("%T\n", c)
	fmt.Println("--------")

	// check channel receive
	chReceive := make(<-chan int)

	go func() {
		// invalid operation: chReceive <- 42 (send to receive-only type <-chan int)
		// chReceive <- 42 // Error
		fmt.Println(<-chReceive)
	}()

	fmt.Printf("%T\n", c)
	fmt.Println("--------")

	// doesn’t assign

	// specific to general
	// Error: cannot use cr (type <-chan int) as type chan int in assignment
	// c = cr
	// c = cs

	// specific to specific
	// Error: cannot use cs (type chan<- int) as type <-chan int in assignment
	// cr = cs

	// assign

	// general to specific
	cr = c
	cs = c
	fmt.Println(cr, cs)

	// conversion
	fmt.Println("-----")
	// general to specific
	fmt.Printf("c\t%T\n", (<-chan int)(c)) // receive
	fmt.Printf("c\t%T\n", (chan<- int)(c)) // send

	// specific to general
	// Error: cannot convert cr (type <-chan int) to type chan int
	// fmt.Printf("c\t%T\n", (chan int)(cr)) // receive
	// fmt.Printf("c\t%T\n", (chan int)(cs)) // send

}
