package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x) // [0 0 0 0 0]
	// assigning
	x[3] = 30
	fmt.Println(x)
	// len
	fmt.Println(len(x))
}
