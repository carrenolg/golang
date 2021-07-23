package main

import (
	"fmt"
)

func main() {
	// channel cancel
	jobs := make(chan int)

	go func() {
		/*for {
			j, more := <-jobs
			//fmt.Println(j, more)
			if more {
				//time.Sleep(1000 * time.Millisecond)
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
			}
		}*/
		for i := 1; i < 5; i++ {
			jobs <- i
			//fmt.Println("send job", i)
			//time.Sleep(10 * time.Millisecond)
		}
		close(jobs)
	}()

	for j := range jobs {
		//fmt.Println("received job", j)
		fmt.Println(j)
	}
	// close(jobs)
	fmt.Println("send all jobs")
}
