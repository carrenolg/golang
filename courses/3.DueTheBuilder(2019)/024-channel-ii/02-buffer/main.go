package main

import (
	"fmt"
	"time"
)

func main() {
	// buffer
	canal := make(chan int, 1)
	fmt.Println("len:", len(canal), "cap:", cap(canal))

	go func() {
		canal <- 1
		fmt.Println("Ya la envié!")
		fmt.Println("len:", len(canal), "cap:", cap(canal))
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Recibí", <-canal)
	fmt.Println("len:", len(canal), "cap:", cap(canal))
}
