package main

import "fmt"

func main() {
	// overflow
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	// unary addition
	var p int
	fmt.Println(+p, -p)

	// binary operatior
	var a byte = 1
	var b byte = 1 << 1
	var c byte = 1<<1 | 1<<5
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println(int(a))
	fmt.Println(int(b))
	fmt.Println(int(c))

	// set theory with bit operator
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // the set {1, 5}
	fmt.Printf("%08b\n", y) // the set {1, 2}

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)
}
