package main

import "fmt"

func main() {
	c := make(chan int)
	//r := make(chan int)

	go send(c)

	receive(c)

}

func send(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receive(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
