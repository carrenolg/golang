package main

import (
	"context"
	"fmt"
	"time"
)

type ctxKey int

var k ctxKey = 0

func procesoLento(ctx context.Context, done chan<- struct{}) {
	select {
	case <-ctx.Done():
		fmt.Println("Cancelado!", ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("Completado!", ctx.Value(k))
	}
	done <- struct{}{}
}

func main() {
	// Context
	fmt.Println("Pkg context")
	comienzo := time.Now()
	done := make(chan struct{})

	ctx := context.WithValue(context.Background(), k, "Hola Context!")

	go procesoLento(ctx, done)

	<-done
	tiempo := time.Since(comienzo)
	fmt.Println("tiempo:", tiempo)
}
