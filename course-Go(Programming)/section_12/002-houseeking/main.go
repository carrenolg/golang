package main

import "fmt"

// underlying type
// IT IS A BAD PRACTICE TO ALIAS TYPES
type foo int

func main() {
	// user-define types
	var x foo
	var y int
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// int convertion
	// x = int(42) // Error: cannot use int(42) (type int) as type foo in assignmentgo
	x = foo(42)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// conert foo to int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// Named types vs. Anonymous types
	var a int = 24 //the compiler hasn't flexibility
	var b = 50     //the compiler has flexibility
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// variable with a certain type
	var z int
	z = 50
	fmt.Println(z)

}
