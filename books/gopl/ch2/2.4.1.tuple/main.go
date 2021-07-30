package main

import "fmt"

func main() {
	// tuple assignment
	v := 1
	fmt.Println(v)
	//fmt.Println(v++) // it's not a expression
	v++ // statement
	fmt.Println(v)

	// GCD (125, 50)
	r := gcd(125, 50)
	fmt.Println("gcd(125, 50):", r)

	// assignability
	var x int
	var s []int
	// x = nil "error"
	x = 0
	s = nil
	fmt.Println(x, s)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
