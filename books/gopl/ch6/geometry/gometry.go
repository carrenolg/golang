package geometry

import (
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type Path []Point // Named type

type ColoredPoint struct {
	Point // embedded field and anonymous field
	Color color.RGBA
}

// Using Pointer
type ColoredPointPointer struct {
	*Point // embedded field and anonymous field
	Color  color.RGBA
}

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

// Add
func (p Point) Add(q Point) Point {
	return Point{
		p.X + q.X,
		p.Y + q.Y,
	}
}

// Sub
func (p Point) Sub(q Point) Point {
	return Point{
		p.X - q.X,
		p.Y - q.Y,
	}
}

// TranslateBy: method of type Path
func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}
