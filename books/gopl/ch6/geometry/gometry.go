package geometry

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	// implementation
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // function call
	fmt.Println(p.Distance(q))  // mathod call
}

/*
Terminology

receiver: use to call to a method
*/
