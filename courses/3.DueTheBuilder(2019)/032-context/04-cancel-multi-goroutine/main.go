package main

import (
	"context"
	"fmt"
	"time"
)

func procesoLento(ctx context.Context, done chan<- struct{}) {
	d2 := make(chan struct{})
	go otroProceso(ctx, d2)
	select {
	case <-ctx.Done(): // get cancel signal from context
		fmt.Println("Cancelado!", ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("Completado!")
	}
	<-d2
	done <- struct{}{}
}

func otroProceso(ctx context.Context, done chan<- struct{}) {
	select {
	case <-ctx.Done(): // get cancel signal from context
		fmt.Println("Completado!", ctx.Err())
	case <-time.After(5 * time.Second):
	}
	done <- struct{}{}
}

func main() {
	// Context
	fmt.Println("Pkg context")
	comienzo := time.Now()
	done := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())

	go procesoLento(ctx, done)

	// cancel goroutine
	cancel()

	<-done
	tiempo := time.Since(comienzo)
	fmt.Println("tiempo:", tiempo)
}
