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
	// the type hotdog is not same that type int
	// error y = x, those are differents types

}
