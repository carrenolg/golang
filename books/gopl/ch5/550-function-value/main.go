package main

import "fmt"

func square(n int) int {
	return n * n
}

func main() {
	// function value
	f := square
	fmt.Println(f(3)) // "9"
}
