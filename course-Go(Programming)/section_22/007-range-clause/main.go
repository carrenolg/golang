package main

import "fmt"

func main() {
	ch := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// channel is closed
		close(ch)
	}()

	// receive
	for v := range ch {
		fmt.Println(v)
	}

	// channel is cloused (you get a value nil or 0)
	fmt.Println("ch close")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
