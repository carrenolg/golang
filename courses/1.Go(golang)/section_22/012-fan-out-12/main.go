package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	// send values
	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		// send values
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	const goroutines = 7

	wg.Add(goroutines)
	var count int = 0
	for i := 0; i < goroutines; i++ {

		// create a new goroutine
		count++
		go func() {
			// receive value
			for v := range c1 {
				func(v2 int) {
					// send value
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			fmt.Println("Gorutines ########:\t", runtime.NumGoroutine())
			wg.Done()
		}()

	}
	wg.Wait()
	close(c2)
	fmt.Println("total goroutine:", count)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
