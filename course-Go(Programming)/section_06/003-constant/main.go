package main

import "fmt"

// untype constant
const a = 15
const b = 20.56988
const c = `Hi, this is Golang`

// type constant

const (
	x int     = 15
	y float64 = 3.14568
	z string  = "Hello world"
)

// iota
const (
	q = iota
	w = iota
	e = iota
)

const (
	r = iota
	t = iota
	u = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	// untype
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	// successive untyped integer constants
	fmt.Println(q)
	fmt.Println(w)
	fmt.Println(e)
	fmt.Printf("%T\n", q)
	fmt.Printf("%T\n", w)
	fmt.Printf("%T\n", e)

	fmt.Println(r)
	fmt.Println(t)
	fmt.Println(u)
	fmt.Printf("%T\n", r)
	fmt.Printf("%T\n", t)
	fmt.Printf("%T\n", u)
}
