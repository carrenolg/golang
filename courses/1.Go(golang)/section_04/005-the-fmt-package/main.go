package main

import "fmt"

// literal value (30)
// the compilador assign the type
var y = 30

func main() {
	// the fmt package
	fmt.Printf("%T\n", y)  // type: int
	fmt.Printf("%#v\n", y) // Go-syntax
	fmt.Println("y:", y)

	// integer:
	fmt.Printf("%b\n", y)  // binary
	fmt.Printf("%o\n", y)  // octal
	fmt.Printf("%x\n", y)  // hexa
	fmt.Printf("%#x\n", y) // hexa with 0 in front
}
