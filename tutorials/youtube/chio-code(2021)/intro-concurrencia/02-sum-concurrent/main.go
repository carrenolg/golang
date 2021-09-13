package main

import (
	"fmt"
)

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum // send result into the channel
}

func main() {
	// channels
	s := []int{0, 1, 2, 3, 4, 5, 6}
	ch := make(chan int)

	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	r1, r2 := <-ch, <-ch // receive result from sum()
	fmt.Println("Total:", r1+r2)
}
