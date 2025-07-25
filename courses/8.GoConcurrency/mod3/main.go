package main

import "fmt"

func main() {
	fmt.Println("Init module 3..")
	c := make(chan int)
	go prod(1, 2, c)
	go prod(3, 4, c)
	a := <-c
	b := <-c
	fmt.Println(a * b)
}

func prod(v1 int, v2 int, c chan int) {
	c <- v1 * v2
}
