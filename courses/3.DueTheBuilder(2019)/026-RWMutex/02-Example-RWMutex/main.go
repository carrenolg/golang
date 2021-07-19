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
	// RWMutex
	var rw sync.RWMutex
	m := make(map[int]int)

	// loop
	for i := 0; i < 100000; i++ {
		go func(i int) {
			for {
				time.Sleep(10 * time.Millisecond)
				rw.RLock()
				x := m[i] // read
				rw.RUnlock()
				fmt.Println(doble(x))
			}
		}(i)
	}

	// loop
	for i := 0; i < 1000; i++ {
		time.Sleep(10 * time.Millisecond)
		rw.Lock()
		m[i] = doble(i) // write
		rw.Unlock()
	}

	fmt.Println(m[100])
}
