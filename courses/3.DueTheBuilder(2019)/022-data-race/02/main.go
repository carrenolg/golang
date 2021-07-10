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

	for i := 0; i < gs; i++ {
		wg.Add(1)
		go func() {
			// read
			v := counter
			runtime.Gosched() // yield thread
			// increment by 1
			v++
			// write
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Count:", counter)
}
