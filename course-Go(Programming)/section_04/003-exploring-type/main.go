package main

import "fmt"

// Declare variable TYPE int
var z int = 10

// declare varible of TYPE string
var y string = "Today is friday"

func main() {
	var x int = 10
	fmt.Println(x)
	x = 6897
	fmt.Println(x)
	// Declare a new variable with the same IDENTIFIER
	// var x string = "hello world"
	// Error: x redeclared in this block

	// the z variable is the type int
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	// the y variable is the type int
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	// Dynamic assignation
	// y = 10
	// cannot use 10 (type untyped int) as type string in assignment

	// raw string literals
	var s string = `Hello world, "name"`
	fmt.Println(s)

	var a = `Hola, my name is Juan
Tel: 315684
email: carreno@hotmail.com`
	fmt.Println(a)
}
