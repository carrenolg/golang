package main

import (
	"fmt"
	"runtime"
)

var x = 10
var y = 3.14587

// int
var a uint8 = 255 // (byte)
var b uint8 = 255 // (byte)

func main() {
	// numeric
	// int = whole numbers
	// decimal = real numbers
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	// int8
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	c := a + b
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	d := 200 + 200
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	// runtime
	fmt.Println(runtime.GOOS)   // linux
	fmt.Println(runtime.GOARCH) // amd64

}
