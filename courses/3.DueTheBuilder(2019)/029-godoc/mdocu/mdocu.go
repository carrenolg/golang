package main

import (
	"fmt"

	"example.com/user/mdocu/fib"
)

func main() {
	// module mdocu
	fmt.Println("Hello world!")
	fmt.Println("fib(5):", fib.FibCiclo(5))
	fmt.Println("fib(10):", fib.FibRecursivo(10))
}
