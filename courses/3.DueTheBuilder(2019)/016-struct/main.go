package main

import "fmt"

// define
type Point struct {
	X float64
	Y float64
}

type Circle struct {
	Center Point // named fields
	Radius float64
}

// embebed struct
type Wheel struct {
	Circle // anony fields
	Ray    int
}

func main() {
	// struct
	var p Point
	fmt.Printf("%#v \n", p)
	p.X = 5.0
	p.Y = 10.0

	q := Point{X: 5.0, Y: 10.0}
	fmt.Printf("%#v \n", q)

	c := Circle{Center: Point{X: 5, Y: 5}, Radius: 5.5}
	fmt.Printf("%#v \n", c)
	fmt.Printf("%#v \n", c.Center.X)

	var w Wheel
	fmt.Printf("%#v \n", w)
	w2 := Wheel{Circle: Circle{Center: Point{X: 5, Y: 5}, Radius: 5.5}, Ray: 15}
	fmt.Printf("%#v \n", w2)
	fmt.Printf("%#v \n", w2.Center.X) // short access to X

	// compare struct
	fmt.Println(p == q) // true

	// slice or fields of type slice (no compare) cannot be compared
}
