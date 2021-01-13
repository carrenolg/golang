package main

// one goroutine is the main
// goroutine that comes by default
import (
	"fmt"
	"runtime"
	"sync"
)

var wgIns sync.WaitGroup

func main() {

	// shared variable
	counter := 0

	// the other 10 goroutines are
	// supposed to come from here
	wgIns.Add(10)

	var mu sync.Mutex

	for i := 0; i < 10; i++ {

		// goroutines are made
		go func() {
			(&mu).Lock()
			for j := 0; j < 10; j++ {
				// shared variable execution
				counter++
				// 100 should be the counter value but
				// it may be equal to 100 or lesser
				// due to race condition
			}
			(&mu).Unlock()
			wgIns.Done()
		}()
	}

	fmt.Println("The number of goroutines before wait = ", runtime.NumGoroutine())

	wgIns.Wait()

	// value should be 100
	fmt.Println("Counter value = ", counter)

	fmt.Println("The number of goroutines after wait = ", runtime.NumGoroutine())

}
