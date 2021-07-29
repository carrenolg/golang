package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	// pointers
	x := 1
	p := &x // p, of type *int, points to x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x) // "2"

	// return address
	var r = f()
	fmt.Println(*r) // "1"
	// fmt.Println(f() == f()) // "false"

	// passing address
	v := 1
	fmt.Printf("v: %d, &v: %v\n", v, &v)
	v = incr(&v)
	fmt.Printf("v: %d, &v: %v\n", v, &v)
	fmt.Println(incr(&v))
	fmt.Printf("v: %d, &v: %v\n", v, &v)
}
