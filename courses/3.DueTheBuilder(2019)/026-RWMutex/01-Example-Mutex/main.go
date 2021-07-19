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
	// Mutex
	var mu sync.Mutex // protege a m
	m := make(map[int]int)

	// loop
	for i := 0; i < 100000; i++ {
		go func(i int) {
			for {
				time.Sleep(10 * time.Millisecond)
				mu.Lock()
				x := m[i] // read
				mu.Unlock()
				fmt.Println(doble(x))
			}
		}(i)
	}

	// loop
	for i := 0; i < 1000; i++ {
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		m[i] = doble(i) // write
		mu.Unlock()
	}

	fmt.Println(m[100])
}
