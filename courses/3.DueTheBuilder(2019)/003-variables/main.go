package main

import "fmt"

// user determines of type
var x int = 1

// compiler determines of type
var y = 1

// no explicit initialization is provided
var z int

// zero value: default value

func main() {
	fmt.Println("The value is:", x)
	fmt.Println("The value is:", y)
	fmt.Println("The value is:", z)

	var a float64 = 1
	var b = 1
	fmt.Printf("Type of a is %T\n", a)
	fmt.Printf("Type of b is %T\n", b)
}
