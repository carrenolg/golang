package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(t1)
	// create 3 channels
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	//go send(eve, odd, quit)
	go send(eve, odd, quit)

	// receive
	//receive(eve, odd, quit)
	receive(eve, odd, quit)

	t2 := time.Now()
	fmt.Println(t2)

	diff := t2.Sub(t1)
	fmt.Println(diff)
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- true
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the eve channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case <-q:
			return
		}
	}
}
