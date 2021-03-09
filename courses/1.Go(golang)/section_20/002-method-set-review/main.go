package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) duplicate() float64 {
	return c.radius + c.radius
}

func main() {
	// method set
	c := &circle{2.5}
	fmt.Println(c.area())
	fmt.Println(c.duplicate())
}
