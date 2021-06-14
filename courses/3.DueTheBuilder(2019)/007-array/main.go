package main

import "fmt"

func main() {
	// array: aggregated data type
	var a [3]int
	for i, v := range a {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// assigment
	a[0] = 0
	a[1] = 1
	a[2] = 2

	for i, v := range a {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// init array
	// literal assigment
	var b = [3]int{1, 2, 3}
	for i, v := range b {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// spread operator ...
	var c = [...]int{1, 2, 3}
	for i, v := range c {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// numbers of element belongs to data type
	// [4]int != [3]int
	var x [3]int
	var y [4]int
	fmt.Printf("x: %T, y: %T\n", x, y)
	var j = [...]int{1, 2, 3}
	var k = [3]int{1, 2, 3}
	fmt.Println("equal:", j == k) // true

	// explicit index
	var ind = [3]int{0: 10, 1: 20}
	for i, v := range ind {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

}
