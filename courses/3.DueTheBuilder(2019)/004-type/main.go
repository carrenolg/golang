package main

import "fmt"

func main() {
	// int
	a, b := 5, 4
	x := a + b
	fmt.Printf("\nVariable x of type: [%T], value: %[1]v\n", x)
	// float
	j, k := 5.0, 4.0
	y := j / k
	fmt.Printf("\nVariable y of type: [%T], value: %[1]v\n", y)
	// working with bits
	// shift operator <<
	n, i := 1, uint(1)
	r := n << i
	fmt.Printf("\nVariable y of type: [%T], value: %[1]v\n", r)
}

// data types

// int int8 int16 int32 int64

// uint uint8 uint16 uint32 uint64

// float32 float64

// bool

// byte rune

// string
