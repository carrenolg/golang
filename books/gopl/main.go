package main

import (
	"fmt"

	"carrenolg.io/books/gopl/ch6/geometry"
)

func main() {
	fmt.Println("6 .Methods")
	fmt.Println("6.1 Methods Declarations")
	// type Point
	p := geometry.Point{
		X: 1,
		Y: 2,
	}
	q := geometry.Point{
		X: 2,
		Y: 4,
	}
	fmt.Println(geometry.Distance(p, q)) // function call
	fmt.Println(p.Distance(q))           // method call

	// type Path
	perim := geometry.Path{
		geometry.Point{
			X: 1,
			Y: 1},
		geometry.Point{
			X: 5,
			Y: 1},
		geometry.Point{
			X: 5,
			Y: 4,
		},
		geometry.Point{
			X: 1,
			Y: 1,
		},
	}
	fmt.Println(geometry.Path.Distance(perim)) // standalone function
	fmt.Println(perim.Distance())              // method of geometry.Path

}
