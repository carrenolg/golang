package main

import "fmt"

type A struct {
	field1 int
	field2 int
}

type B struct {
	A
	field3 int
	field4 int
}

func main() {
	// promoved fields
	v := B{
		A:      A{1, 2},
		field3: 3,
		field4: 4,
	}
	fmt.Printf("%T, %#[1]v\n", v)
	fmt.Println(v.field1, v.field2, v.field3, v.field4)

	// promoved methods
}
