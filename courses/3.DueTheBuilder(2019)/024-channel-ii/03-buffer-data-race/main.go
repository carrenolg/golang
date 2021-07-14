package main

import (
	"fmt"
	"runtime"
	"time"
)

type CarMarket struct {
	subtotal float64
	tax      float64
	shipping float64
	total    float64
}

func calcCargos(entrada <-chan *CarMarket, salida chan<- *CarMarket) {
	// receiving data from carritos
	for c := range entrada {
		time.Sleep(3 * time.Second)
		c.tax = c.subtotal * 0.075
		c.shipping = c.subtotal * 0.10
		// sending data to conCargos
		salida <- c
	}
	close(salida)
}

func calcTotal(entrada <-chan *CarMarket, salida chan<- *CarMarket) {
	// receving data from conCargos
	for c := range entrada {
		time.Sleep(3 * time.Second)
		c.total = c.subtotal + c.tax + c.shipping
		// sending data to conTotal
		salida <- c
	}
	close(salida)
}

func main() {
	// buffer for resolve the data race condition
	subtotales := []float64{100.00, 110.00, 50.00, 55.00}

	carritos := make(chan *CarMarket, len(subtotales))
	conCargos := make(chan *CarMarket, len(subtotales))
	conTotal := make(chan *CarMarket, len(subtotales))

	go calcCargos(carritos, conCargos)
	go calcTotal(conCargos, conTotal)

	// sending data to carritos
	for _, st := range subtotales {
		carritos <- &CarMarket{subtotal: st}
	}
	close(carritos)

	// reciving data from conTotal
	for c := range conTotal {
		fmt.Printf("\n %#v \n", c)
		fmt.Println("Goroutine#:", runtime.NumGoroutine())
	}
}
