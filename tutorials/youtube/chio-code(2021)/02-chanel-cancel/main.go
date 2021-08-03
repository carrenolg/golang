package main

import (
	"fmt"
	"time"
)

func main() {
	// channel cancel
	jobs := make(chan int)
	// done := make(chan bool)

	go func() {
		for i := 1; i < 5; i++ {
			jobs <- i
			fmt.Println("send job", i)
			time.Sleep(10 * time.Millisecond)
		}
		close(jobs)
	}()

	for j := range jobs {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("received job", j)
	}
	// not working
	/*for {
		j, more := <-jobs
		//fmt.Println(j, more)
		if more {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("received job", j)
		} else {
			fmt.Println("received all jobs")
			done <- true
			close(done)
			break
		}
	}*/
	fmt.Println("send all jobs")
}
