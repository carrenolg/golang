package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var m map[int]int

func mapInit() {
	m = make(map[int]int)
	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Millisecond)
		m[i] = i * 2
	}
}

func entrada(k int) int {
	mu.RLock()
	if m != nil {
		e := m[k]
		mu.RUnlock()
		return e
	}
	mu.RLock()
	if m == nil {
		mapInit()
	}
	e := m[k]
	mu.RUnlock()

	return e
}

func main() {
	// WRMutex
	// 1s == 1.000.000.000n (1e+9)
	comienzo := time.Now()
	x := entrada(100)
	fmt.Printf("sin inicializar: %d %v\n", x, time.Since(comienzo))

	comienzo = time.Now()
	fmt.Printf("inicializado: %d, %v\n", x, time.Since(comienzo))
}
