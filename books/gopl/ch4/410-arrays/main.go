package main

import "fmt"

func zero(arr [3]rune) {
	for i := range arr {
		arr[i] = 0
	}
	fmt.Println(arr)
}

func zeroPtr(ptr *[3]rune) {
	*ptr = [3]rune{}
	fmt.Println(*ptr)
}

func main() {
	// 4.1 Arrays
	// ellipsis "..."
	a := [...]int{1, 2, 3}
	fmt.Println(a) // "[1 2 3]"

	// explicit index
	b := [3]int{0: 1, 1: 2, 2: 3}
	fmt.Println(b) // "[1 2 3]"

	// index in any order
	c := [...]int{5: -1, 4: 0}
	fmt.Println(c) // "[0 0 0 0 0 -1]"

	// comparable arrays
	d := [2]int{1, 2}
	e := [3]int{1, 2}
	f := [3]int{1, 2}
	fmt.Println(e == f, d) // "true [1 2]"
	// >> fmt.Println(d == e) // compile error: cannot compare

	// arrays like arguments
	g := [3]rune{'a', 'b', 'c'}
	// inefficient, passing by value
	zero(g) // "[0 0 0]"
	g = [3]rune{'a', 'b', 'c'}
	// efficient, passing by reference
	zeroPtr(&g) // "[0 0 0]"
}
