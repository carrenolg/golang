package main

import "fmt"

func main() {
	// no variadic, fixed parameters
	foo(1, 2, 3)

	// variadic
	// passing arguments
	variadicFoo(0, 1, 2, 3, 4, 5)
	variadicFoo(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	variadicFoo(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	variadicFoo(0, 1, 2, 3, 4)

	x := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Printf("%T\n", x) // slice: []int

	// passing arguments
	// lexical symbol [x...]
	// this operator take a slice of int to convert to individual element
	variadicFoo(x...)

	unlimitedParameter("unlimited", 0, 1, 2, 3, 4, 5)

}

func foo(x, y, z int) {
	fmt.Println(x, y, z)
}

// lexical symbol [...int] (triple dots)
func variadicFoo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // slice: []int
}

// final parameter
func unlimitedParameter(s string, x ...int) {
	fmt.Println(s)
	fmt.Println(x)
}
