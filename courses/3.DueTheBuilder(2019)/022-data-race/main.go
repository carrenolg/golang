package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type CarMarket struct {
	subtotal float64
	tax      float64
	shipping float64
	total    float64
}

func main() {
	// data race
	subtotales := []float64{100.00, 110.00, 50.00, 55.00}

	for _, st := range subtotales {
		var wg sync.WaitGroup
		c := CarMarket{subtotal: st}

		// secuential executetion
		wg.Add(1)
		go calcCargos(&c, &wg)
		fmt.Println("Goroutines running \t", runtime.NumGoroutine())
		//go calcTotal(&c, &wg)

		wg.Wait()
		fmt.Printf("\n %#v \n", c)
	}

}

func calcCargos(c *CarMarket, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	c.tax = c.subtotal * 0.075
	c.shipping = c.subtotal * 0.10
	calcTotal(c)
}

func calcTotal(c *CarMarket) {
	// defer wg.Done()
	time.Sleep(3 * time.Second)
	c.total = c.subtotal + c.tax + c.shipping
}
