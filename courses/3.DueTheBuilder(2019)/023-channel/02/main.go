package main

import (
	"fmt"
	"runtime"
	"time"
)

func genera(salida chan<- int) {
	for i := 0; i < 5; i++ {
		salida <- i
	}
	close(salida)
}

func cuadrado(entrada <-chan int, salida chan<- int) {
	for i := range entrada {
		time.Sleep(500 * time.Millisecond)
		salida <- i * i
	}
	close(salida)
}

func main() {
	// channels
	num := make(chan int)
	cuadrados := make(chan int)
	cuartas := make(chan int)

	// pipelines
	go genera(num)
	go cuadrado(num, cuadrados)
	go cuadrado(cuadrados, cuartas)

	for i := range cuartas {
		fmt.Printf("%d ", i)
		fmt.Println("Goroutine#:", runtime.NumGoroutine())
	}
	fmt.Println()
}
