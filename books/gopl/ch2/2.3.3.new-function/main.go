package main

import "fmt"

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

func main() {
	// how create variables
	p := new(int)   // p, of type int, points to an unnamed int variable
	fmt.Println(*p) // "0"
	*p = 2          // set the unnamed int to 2
	fmt.Println(*p) // "2"

	p1 := newInt()
	p2 := newInt2()
	fmt.Println(p1, p2)
	equal := p1 == p2
	fmt.Println(equal) // "false"
}
