package main

import "fmt"

// the zero value is setting to all variables
var y string
var x int
var a []int

func main() {
	// var y
	fmt.Printf("%T\n", y)
	fmt.Println("y:", y)
	// var x
	fmt.Printf("%T\n", x)
	fmt.Println("x:", x)
	// var a
	fmt.Printf("%T\n", a)
	fmt.Println("a:", a)
}
