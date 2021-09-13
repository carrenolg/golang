package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Lift off!")
}

func main() {
	// gorputine leak
	abort := make(chan struct{})
	tick := make(chan time.Time)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	// select
	fmt.Println("Commencing countdown. Press return to abort.")

	//tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)

		go func() {
			time.Sleep(1 * time.Second)
			tick <- time.Now()
		}()

		select {
		case <-tick:
			// do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}

	launch()
}
