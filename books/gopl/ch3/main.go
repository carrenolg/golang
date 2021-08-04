package main

import "fmt"

func main() {
	// overflow
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+2, i*i)
	fmt.Printf("%08b\n", i+2)

	var j uint8 = 255
	var k uint8 = 25
	fmt.Println(j * k)

	// unary addition
	var p int
	fmt.Println(+p, -p)

	// binary operatior
	var a uint8 = 1
	var b uint8 = 1 << 1
	var c uint8 = 1<<1 | 1<<5
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

	// conversions
	var n = 1e100
	fmt.Printf("%T, %v\n", n, n)
	nint := int(n)
	fmt.Printf("%T, %v\n", nint, nint)

	// write and print integers
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	// hex
	h := int(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", h)
	// rune
	ascii := 'a'
	unicode := 'ㅏ'
	newline := '\n'
	fmt.Printf("%d %[1]c %#[1]q\n", ascii)
	fmt.Printf("%d %[1]c %#[1]q\n", unicode)
	fmt.Printf("%d %#[1]q\n", newline)
}
