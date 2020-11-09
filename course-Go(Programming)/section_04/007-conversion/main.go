package main

import "fmt"

var y int

// create own type (hotdog)
type hotdog int

var x hotdog

func main() {
	y = 25
	x = 20 // hotdog type
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// conversion x of type hotdog into type int
	y = int(x) // conversion
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
