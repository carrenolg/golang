package main

import "fmt"

// package scope
var y = 5

// main scope
func main() {
	// sub-scope
	{
		x := 1
		fmt.Println(x)
	}
	fmt.Println(y)

	// pointers
	// get memory address
	p := &y
	fmt.Println(p)

	// get value (dereference)
	fmt.Println(*p)

	// assign value
	*p = 10
	fmt.Println(*p)

	// inpect pointer
	fmt.Println(y, &y, p, *p)
	fmt.Printf("%T\n", p)

	// Lab
	var q *int
	fmt.Println(q) // nil

}

// notes
// 1. varible: is a memory location
