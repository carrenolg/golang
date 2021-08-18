package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type Point struct{ X, Y int }

// embedding
type Circle struct {
	Point
	Radius int
}

type Whell struct {
	Circle
	Spokes int
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func main() {
	// 4.4 Struct
	// def: a struct is an aggregate data type
	// declare

	// init
	var dilbert Employee
	dilbert.Salary -= 500
	fmt.Println(dilbert) // "{0   0001-01-01 00:00:00 +0000 UTC  -500 0}"

	// accesing by pointer
	position := &dilbert.Position
	*position = "Senior Dev Go" + *position
	fmt.Println(dilbert.Position) // "Senior Dev Go"

	// empty struct, no fields
	seen := make(map[string]struct{})
	if _, ok := seen["s"]; !ok {
		seen["s"] = struct{}{}
	}
	fmt.Println(seen) // "map[s:{}]"

	// 4.4.1 Struct Literals
	p := Point{1, 2}
	fmt.Println(p) // "{1 2}"
	q := Point{X: 1, Y: 2}
	fmt.Println(q) // "{1 2}"

	// passing structures like arguments
	// value
	fmt.Println(Scale(Point{1, 2}, 5)) // "{5 10}"
	// pointer
	pp := &Employee{Salary: 25}
	fmt.Println(Bonus(pp, 10)) // "2"

	//comparing structs
	fmt.Println(q == p)                   // "true"
	fmt.Println(p.X == q.X && p.Y == q.Y) // "true"

	// struct embedding
	var w Whell
	w.X = 8      // equivalent to w.Circle.Point.X = 8
	w.Y = 8      // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20
	// note: struct literals is not support with shorhand

}
