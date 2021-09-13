package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("Lift off!")
}

func main() {
	// lab
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println("countdown:")
		<-tick // loop iteration wait until time.Tick send the value
	}
	// launch
	launch()
}
