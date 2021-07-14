package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	// data race simulated (yield thread)
	counter := 0

	var wg sync.WaitGroup
	const gs = 100
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		wg.Add(1)
		go func() {
			mu.Lock() // blocking variable
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() // unbloking variable
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Count:", counter)
}
