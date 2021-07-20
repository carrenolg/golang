package main

import (
	"fmt"
	"sync"
	"time"
)

func doble(n int) int {
	return n * 2
}

func main() {
	// sync.Map
	m := new(sync.Map)

	// loop
	for i := 0; i < 100000; i++ {
		go func(i int) {
			for {
				time.Sleep(10 * time.Millisecond)
				x, ok := m.Load(i) // read
				if !ok {
					x = 0
				}
				fmt.Println(doble(x.(int))) // type assertion
			}
		}(i)
	}

	// loop
	for i := 0; i < 1000; i++ {
		time.Sleep(10 * time.Millisecond)
		m.Store(i, doble(i)) // write
	}

	fmt.Println(m.Load(100))
}
