package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"carrenolg.io/books/gopl/ch6/intset"
)

func main() {
	// 7.3 Interface Satisfaction
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	// types "os.Stdout" and "*bytes.Buffer" satisfies the io.Writer interface
	// w = time.Second // compile error, time.Duration lacks Write method
	fmt.Printf("%T\n", w)

	// type have a method
	// var _ = intset.IntSet{}.String(), // compile error

	var s intset.IntSet
	var _ = s.String() // *IntSet: has a String methos, but IntSet hans't

	var _ fmt.Stringer = &s // ok
	// var _ fmt.Stringer = s // compile error

	// empty interface type
	var any interface{}
	any = true
	fmt.Printf("%T\n", any)
	any = 12.34
	fmt.Printf("%T\n", any)

	// asserts at compile time that a value of type
	// *bytes.Buffer satisfy io.Writer
	var wr io.Writer = new(bytes.Buffer)
	fmt.Printf("%T\n", wr)
}
