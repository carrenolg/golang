package main

import (
	"fmt"
	"math"
)

// create a new type
type circle struct {
	radius float32
}

// create a interface
type shape interface {
	area() float32
}

// attache method to circle type
// pointer receiver
func (c *circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

// non-pointer receiver
/*func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}*/

// create func to implement the interface
func info(s shape) {
	fmt.Println("Area:", s.area())
}

func main() {

	c := circle{
		radius: 2.5,
	}
	// non-pointer value
	// info(c)

	// pointer value
	info(&c)

}
