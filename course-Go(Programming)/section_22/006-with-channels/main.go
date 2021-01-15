package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(t1)
	// create 3 channels
	var eve []int
	var odd []int

	// send
	//go send(eve, odd, quit)
	e, o := send(eve, odd)

	fmt.Println(e)
	fmt.Println(o)

	// receive
	//receive(eve, odd, quit)
	receive(e, o)

	t2 := time.Now()
	fmt.Println(t2)

	diff := t2.Sub(t1)
	fmt.Println(diff)
}

func send(e, o []int) ([]int, []int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e = append(e, i)
		} else {
			o = append(o, i)
		}
	}
	return e, o
}

func receive(e, o []int) {
	for i := range e {
		fmt.Println("from the eve slice:", i)
	}
	for i := range o {
		fmt.Println("from the odd slice:", i)
	}
}
