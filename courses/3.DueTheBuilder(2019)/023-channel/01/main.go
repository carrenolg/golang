package main

import (
	"fmt"
	"time"
)

func doble(n int) int {
	return n * 2
}

func main() {
	// channels
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- doble(i) // send
		}
		// deadlock: channel open
		close(ch) // channel close: return 0
	}()

	// idiomatic
	// range knows implicitly if the channel is closed
	for i := range ch {
		fmt.Println("i:", i)
	}

	// no idiomatic
	// you have to check the channel explicitly
	for {
		i, ok := <-ch // recive
		// i: send value
		// ok: reports if channel is close or open
		if !ok {
			break
		}
		fmt.Println(i)
	}

}
