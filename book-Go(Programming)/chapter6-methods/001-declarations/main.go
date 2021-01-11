package main

import (
	"fmt"
	"math"
)

type point struct{ x, y float64 }

type path []point

// traditional function
func distance(p, q point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// method of the point type
func (p point) distance(q point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// method of the path type
func (p path) distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].distance(p[i])
		}
	}
	return sum
}

// method of the point type (with a pointer receiver)
func (p *point) scaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func main() {
	// implement
	p := point{1, 2}
	q := point{4, 6}

	fmt.Println(distance(p, q))
	fmt.Println(p.distance(q)) // method call

	// implement
	perim := path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.distance())

	// implement (pointer receiver)
	p2 := point{1, 2}
	p2.scaleBy(2)
	// the compailer will perform an implicit &p
	fmt.Println(p2) // shorthand

	// call with pointer
	pptr := &point{1, 2}
	fmt.Printf("%T\n", pptr)
	fmt.Println(pptr.distance(q))

}
