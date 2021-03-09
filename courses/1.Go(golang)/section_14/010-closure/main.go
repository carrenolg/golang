package main

import "fmt"

// playing with the scope
// the scope of x is this whole package
var x int

func main() {
	// playing
	fmt.Println(x)
	x++
	// closure
	{
		var y int = 20
		fmt.Println(y)
		y++
		fmt.Println(y)
	}
	// fmt.Println(y) Error: undefined
	fmt.Println(x)
	foo()

	// closure
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	// playing : wowwww
	c := b
	fmt.Println(c())
}

func foo() {
	fmt.Println("hello from foo")
	x++
	fmt.Println(x)
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
