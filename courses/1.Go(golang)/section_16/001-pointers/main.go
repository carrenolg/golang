package main

import "fmt"

func main() {
	// pointers
	a := 10
	fmt.Println(a)

	// get the memory address and create the pointer p
	p := &a
	fmt.Println(p)
	fmt.Printf("%T\n", p) // type: [*int]

	// get the value stored in memory address
	v := *p
	fmt.Println(v)
	fmt.Printf("%T\n", v) // type: [*int]

	// change the value which the pointer points
	*p = 11
	fmt.Println(a)
	fmt.Printf("%T\n", a) // type: [*int]

}
