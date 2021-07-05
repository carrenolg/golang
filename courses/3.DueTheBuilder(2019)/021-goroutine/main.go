package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// info
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU's\t", runtime.NumCPU())
	//fmt.Println("Goroutine\t", runtime.NumGoroutine())
	var wg sync.WaitGroup
	// goroutine
	op := func(i int) {
		defer wg.Done()
		fmt.Println("op", i, time.Now().Format("5"))
		time.Sleep(3 * time.Second)
		fmt.Println("op", i, time.Now().Format("5"))
	}
	start := time.Now()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go op(i)
		fmt.Println("Goroutine \t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Total:", time.Since(start))
}
