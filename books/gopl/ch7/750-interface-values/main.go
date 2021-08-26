package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	// 7.5 Interface Values
	var w io.Writer
	// Case 1, interface value is nil
	fmt.Printf("type: %T, value: %[1]v\n", w)
	// w.Write([]byte("hello")) // panic: nil pointer dereference

	// Case 2, interface value is *os.File
	w = os.Stdout // implicit conversion io.Writer(os.Stdout)
	fmt.Printf("type: %T, value: %[1]v\n", w)
	w.Write([]byte("hello")) // writes bytes to standard output
	fmt.Printf("\n")

	// Case 3, interface value is *bytes.Buffer
	w = new(bytes.Buffer)
	fmt.Printf("type: %T, value: %[1]v\n", w) // w is not a variable, it's a pointer
	w.Write([]byte("hello"))                  // writes "hello" to a memory buffer

	// Case 4, interface value is nil again
	// check Case 1, both the case 1 and 4 have the same behavior

	fmt.Println("\nComparing Interface values")
	var x interface{} = time.Now()
	y := time.Now()
	// x is a inetrface value
	fmt.Println(x == y) // "false"

	// comparing
	var value1, value2 io.Writer
	value1, value2 = os.Stdout, os.Stdout
	// dynamic type and dynamic value are equal
	fmt.Println(value1 == value2) // "true"
	value2 = new(bytes.Buffer)
	// dynamic type and dynamic value are not equal
	fmt.Println(value1 == value2) // "false"
}
