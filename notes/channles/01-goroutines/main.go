package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Lanzando goroutine...")
	// parallel
	go cincoVeces("Esta goroutine no simpre se completara.")
	// secuencial
	cincoVeces("Este msg se mostrata exactamente 5 veces.")
	// sleep
	time.Sleep(5 * time.Second)
	fmt.Println("Finalizando programa.")
}

func cincoVeces(msg string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("(%d de 5) %s\n", i, msg)
	}
}
