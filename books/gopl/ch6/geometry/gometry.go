package geometry

import (
	"math"
)

type Point struct{ X, Y float64 }

type Path []Point // Named type

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance: method of type Path
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// ScaleBy: method of type Point
// Method name: (*Point).ScaleBy
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
