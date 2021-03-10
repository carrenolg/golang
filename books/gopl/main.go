package main

import (
	"fmt"

	"carrenolg.io/books/gopl/ch6/geometry"
)

func main() {
	fmt.Println("hello world")
	p := geometry.Point{
		X: 1,
		Y: 2,
	}
	q := geometry.Point{
		X: 2,
		Y: 4,
	}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
}
