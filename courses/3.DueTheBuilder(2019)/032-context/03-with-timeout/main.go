package main

import (
	"context"
	"fmt"
	"time"
)

func procesoLento(ctx context.Context, done chan<- struct{}) {
	select {
	case <-ctx.Done():
		fmt.Println("Cancelado!", ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("Completado!")
	}
	done <- struct{}{}
}

func main() {
	// Context
	fmt.Println("Pkg context")
	comienzo := time.Now()
	done := make(chan struct{})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	go procesoLento(ctx, done)

	// cancel goroutine
	defer cancel()

	<-done
	tiempo := time.Since(comienzo)
	fmt.Println("tiempo:", tiempo)
}
