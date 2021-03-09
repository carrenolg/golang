package main

import "fmt"

func main() {
	// expression
	f := func() {
		fmt.Println("this is a func expresion")
	}

	g := func(n int) {
		fmt.Println("You sent the year", n)
	}

	// called
	f()

	g(2020)

}
