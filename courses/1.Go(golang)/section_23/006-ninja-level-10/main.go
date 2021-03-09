package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	c := make(chan int)

	go send(c)

	receive(c)
}

func send(c chan int) {
	var wg sync.WaitGroup
	const goroutine = 10
	wg.Add(goroutine)
	for i := 0; i < goroutine; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			wg.Done()
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	}
	wg.Wait()
	close(c)
}

func receive(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
