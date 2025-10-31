package main

import "sync"

// global variables
var i int = 0
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	println(i) // expected: 2, actual: ?
}

func inc() {
	i = i + 1
	wg.Done()
}
