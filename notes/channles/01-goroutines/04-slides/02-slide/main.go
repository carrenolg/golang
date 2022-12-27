package main

import "fmt"

func main() {
	tareasParalelas := 4
	v := []int{0, 1, 3, 1, 0, 7, 8, 9, 3, 3, 0, 2}
	ch := make(chan int)
	// send
	for t := 0; t < tareasParalelas; t++ {
		s := t
		inicio := s * len(v) / tareasParalelas
		fin := (s + 1) * len(v) / tareasParalelas
		// create new task parallel
		go sum(v[inicio:fin], ch)
	}
	// receive
	totalSuma := 0
	for t := 0; t < tareasParalelas; t++ {
		partialSuma := <-ch
		fmt.Println("Receiving:", partialSuma)
		totalSuma += partialSuma
	}
	fmt.Println("total Suma:", totalSuma)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
