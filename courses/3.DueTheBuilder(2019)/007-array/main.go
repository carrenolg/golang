package main

import "fmt"

func main() {
	// array: aggregated data type
	var a [3]int
	for i, v := range a {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// assigment
	a[0] = 0
	a[1] = 1
	a[2] = 2

	for i, v := range a {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// init array
	// literal assigment
	var b = [3]int{1, 2, 3}
	for i, v := range b {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// spread operator ...
	var c = [...]int{1, 2, 3}
	for i, v := range c {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// numbers of element belongs to data type
	// [4]int != [3]int
	var x [3]int
	var y [4]int
	fmt.Printf("x: %T, y: %T\n", x, y)
	var j = [...]int{1, 2, 3}
	var k = [3]int{1, 2, 3}
	fmt.Println("equal:", j == k) // true

	// explicit index
	var ind = [3]int{0: 10, 1: 20}
	for i, v := range ind {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// max index
	var maxInd = [...]int{9: 90}
	for i, v := range maxInd {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// Array is a efficient data structure
	// In the memory its elements are allocate in sequence
	// one by one

	var p = [100]int{99: 100}
	// get last item, len function
	fmt.Println("last item:", p[len(p)-1])

	// types: int, string, bool, float32
	var s = [3]string{"a", "b", "c"}
	for i, v := range s {
		fmt.Printf("i: %d v: %v\n", i, v)
	}
	var f = [3]float32{1.1, 2.2, 3.3}
	for i, v := range f {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// Laboratory
	var r = [5]rune{'a', 2, 0xf, 4, '5'} // char, int32, hex
	fmt.Println("r:", r)
	// check
	fmt.Printf("r[0] type: %T\n", r[0]) // int32 ASCII code

	// lab 2
	const index = 10
	var sp = [...]int{index: 100}
	for i, v := range sp {
		fmt.Printf("i: %d v: %v\n", i, v)
	}

	// lab 3
	var var1 = 5
	var var2 = 10
	var sp2 = [...]int{var1, var2}
	for i, v := range sp2 {
		fmt.Printf("i: %d v: %v\n", i, v)
	}
}
