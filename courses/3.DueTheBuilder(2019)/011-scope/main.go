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

	// scope
	var rpointer *int
	{
		r := 10
		fmt.Println("r in scope", r)
		// get address
		rpointer = &r
	}
	fmt.Println("r out scope", *rpointer)
	// Lab
	var q *int
	fmt.Println(q) // nil

}

// notes
// 1. varible: is a memory location
// 2. lifetime: The lifetime of a variable is the location (i.e., place) (not a Typo !) where the variable exists
// 3. scope: The scope of a variable is the locations/places/range in the program
// where the variable is accessible/visible
// 4. universe scope - len cap iota
// 5. impor scope - file level
// 6. local scope - func, loop, blocks, etc
// 7. package scope - main, global variables, etc
