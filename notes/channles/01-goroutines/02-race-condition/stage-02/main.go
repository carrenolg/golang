package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	tareasParalelas := runtime.GOMAXPROCS(0)
	v := []int{0, 1, 3, 1, 0, 7, 8, 9, 3, 3, 0, 2}
	// sync
	// mt := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(tareasParalelas)
	// sum
	totalSuma := int64(0)
	for t := 0; t < tareasParalelas; t++ {
		s := t
		go func() {
			defer wg.Done()
			inicio := s * len(v) / tareasParalelas
			fin := (s + 1) * len(v) / tareasParalelas
			suma := Suma(v[inicio:fin])
			// mutual exclution
			//mt.Lock()
			atomic.AddInt64(&totalSuma, int64(suma))
			//mt.Unlock()
		}()
	}
	// wait
	wg.Wait()
	// output
	fmt.Println("suma", totalSuma)
}

func Suma(porcion []int) int {
	total := 0
	for _, n := range porcion {
		total += n
	}
	return total
}
