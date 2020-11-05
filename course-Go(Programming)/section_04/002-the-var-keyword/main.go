package main

import "fmt"

var y string
var w = "this is a global variable"

func main() {
	// Declare the 'z' variable with type int
	var z int
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	// Declare + assign = Initialization
	x := 10
	fmt.Println(x)
	// print global variable, value = ""
	fmt.Printf("%T\n", y)
	// print global variable
	fmt.Printf("%T\n", w)
	fmt.Printf("%s\n", w)
}

func other() {
	// global variable
	fmt.Println(w)
	// local variable
	// fmt.Println(z) (Erro)
}
