package main

import (
	"fmt"
	"sync"
	"time"
)

var initOnce sync.Once

// m declared, but not init
var m map[int]int

func mapInit() {
	m = make(map[int]int) // init m
	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Millisecond)
		m[i] = i * 2
	}
}

func entrada(k int) int {
	initOnce.Do(mapInit)
	return m[k]
}

func main() {
	// 1s == 1.000.000.000n (1e+9)
	comienzo := time.Now()
	x := entrada(100)
	fmt.Printf("sin inicializar: %d %v\n", x, time.Since(comienzo))

	comienzo = time.Now()
	fmt.Printf("inicializado: %d, %v\n", x, time.Since(comienzo))
}
