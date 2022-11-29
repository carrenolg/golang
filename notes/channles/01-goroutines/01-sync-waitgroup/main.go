package main

import (
	"fmt"
	"sync"
)

func main() {
	const numTareas = 3
	// sync
	wg := sync.WaitGroup{}
	wg.Add(numTareas)
	// run goroutines
	for i := 0; i < numTareas; i++ {
		numTarea := i
		go func() {
			defer wg.Done()
			fmt.Println("Ejecutando Tarea ", numTarea)
		}()
	}
	// wait
	wg.Wait()
	fmt.Println("Completadas todas las tareas. Finalizado")
}
