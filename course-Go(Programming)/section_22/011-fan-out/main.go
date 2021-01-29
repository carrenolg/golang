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

	fmt.Println("CPUs\t\t", runtime.NumCPU())

	// send channel
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
	// we have two bi-directional channels c1 and c2
	var wg sync.WaitGroup

	// receive values
	var count int = 0
	for v := range c1 {
		wg.Add(1)

		// create a new goroutine
		count++
		go func(v2 int) {
			// send values
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
		fmt.Println("Gorutines ########:\t", runtime.NumGoroutine())
	}
	// wait gotoutines is done
	wg.Wait()
	close(c2)
	fmt.Println("total goroutine:", count)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
