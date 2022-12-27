package main

import (
	"fmt"
	"time"
)

func main() {
	dulces := make(chan string, 10)
	// receive
	go engullidor("Marcos", dulces)
	go engullidor("Aina", dulces)
	go engullidor("Judit", dulces)
	// send
	dulces <- "Donut"
	time.Sleep(time.Second)
	dulces <- "Cruasan"
	time.Sleep(time.Second)
	dulces <- "Ensaimada"
	time.Sleep(time.Second)
}

func engullidor(nombre string, dulces <-chan string) {
	for dulce := range dulces {
		fmt.Println(nombre, "come", dulce)
	}
}
