package main

import "fmt"

// scope package level
var a int = 42
var y string = "James Bond"
var z bool = true

// own type
type amoung int

var w int

func main() {
	// print all of the VALUES to one single string
	s := fmt.Sprintf("%d\t%s\t%t", a, y, z)
	fmt.Println(s)

	// create own type
	var x amoung
	fmt.Println(x)        // zero value
	fmt.Printf("%T\n", x) // type: amoung
	x = 42
	fmt.Println(x)

	// conversion
	w = int(x)
	fmt.Println(w)
	fmt.Printf("%T\n", w)
}
