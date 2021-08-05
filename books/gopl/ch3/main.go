package main

import (
	"fmt"
	"math"
)

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

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
	unicode := 'ㅂ'
	newline := '\n'
	fmt.Printf("%d %[1]c %#[1]q\n", ascii)
	fmt.Printf("%d %[1]c %#[1]q\n", unicode)
	fmt.Printf("%d %#[1]q\n", newline)

	// float
	var f float32 = 16777216
	fmt.Println(f == f+1)
	fmt.Println(1 << 24)

	// sci notation
	const Planck = 6.62606957e23
	fmt.Println(Planck)

	// verb
	for i := 0; i < 8; i++ {
		fmt.Printf("x = %d e^x = %.3f\n", i, math.Exp(float64(i)))
	}

	// special values
	var z float64
	fmt.Println(z, -z, 1/z, -1/z)

	// complex numbers
	var c1 complex128 = complex(1, 2)
	var c2 complex128 = complex(3, 4)
	fmt.Println(c1 * c2)
	fmt.Println(real(c1 * c2))
	fmt.Println(imag(c1 * c2))

	// strings
	str := "español"      // letras 7
	fmt.Println(len(str)) // "8"
	fmt.Println(str[0], str[7])

	// sub-string
	fmt.Println(str[0:2]) // "es"

	// string literals
	lit := "hello, world"
	fmt.Println(lit)

	// scape sequences
	fmt.Println("alert", "\a")

	// raw literal
	raw := `Hello
	world`
	fmt.Println(raw)

	// unicode
	symbol := '世'
	code := '\u4e16'
	code2 := '\U00004e16'
	fmt.Println(symbol) // "19990"
	fmt.Println(code)   // "19990"
	fmt.Println(code2)  // "19990"

	fmt.Printf("%b\n", code2) // "100111000010110"

	// A
	letter_a := '\x41'                   // hexa
	fmt.Println(letter_a)                // "65" int32
	fmt.Printf("%b\n", letter_a)         // "1000001" byte
	fmt.Printf("%s\n", string(letter_a)) // "A" string

	// check
	fmt.Println(HasPrefix("hola!", "ola"))
}
