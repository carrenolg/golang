package main

import (
	"fmt"
	"time"
)

func procesoLento(done chan<- struct{}) {
	time.Sleep(5 * time.Second)
	fmt.Println("Completado!")
	done <- struct{}{}
}

func main() {
	// Context
	fmt.Println("Pkg context")
	comienzo := time.Now()
	done := make(chan struct{})

	go procesoLento(done)
	<-done
	tiempo := time.Since(comienzo)
	fmt.Println("tiempo:", tiempo)
}
