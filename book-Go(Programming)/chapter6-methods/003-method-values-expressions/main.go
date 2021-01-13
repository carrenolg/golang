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
	// select a method
	// call a method
	p := point{1, 2}
	q := point{4, 6}

	distanceFromP := p.distance // method value

	// call a method using method value
	fmt.Println(distanceFromP(q))

	var origin point

	fmt.Println(distanceFromP(origin))

}
