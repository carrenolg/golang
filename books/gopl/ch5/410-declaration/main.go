package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func add(x int, y int) int {
	return x + y
}

func zero(x, y int) int {
	return x * y * 0
}

func main() {
	// 4.1 Function Declarations
	fmt.Println(hypot(3, 4)) // "5"
	// I'm a caller
	fmt.Println(add(3, 4)) // "7"

	// type of a func (signature)
	fmt.Printf("%T\n", hypot) // "func(float64, float64) float64"
	fmt.Printf("%T\n", add)   // "func(int, int) int"

	// compating two functions
	// fmt.Println(add == hypot), compile error: mismatch type

	// passing arguments by value
	fmt.Println(zero(5, 5)) // "0"

}
