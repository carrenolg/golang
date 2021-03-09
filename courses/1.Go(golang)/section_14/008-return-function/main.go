package main

import "fmt"

func main() {
	// method 1
	// retunr a function that it returns a float64
	f := foo()
	fmt.Printf("%T\n", f) // type: [func() float64]
	// execute the function and getting the return
	pi := f()
	fmt.Println(pi)
	// method 2
	fmt.Println(foo()())
}

func foo() func() float64 {
	pi := func() float64 {
		return 3.1458
	}
	return pi
}
