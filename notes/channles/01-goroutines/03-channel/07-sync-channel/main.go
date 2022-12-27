package main

import "fmt"

func main() {
	espera := TaskAsync()
	// wait
	<-espera
	fmt.Println("programa finalizado.")
}

func TaskAsync() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		fmt.Println("haciendo algo en paralelo...")
		for i := 0; i < 3; i++ {
			fmt.Println(i, "...")
		}
		fmt.Println("finalizada tarea en paralelo")
		close(ch)
	}()
	return ch
}
