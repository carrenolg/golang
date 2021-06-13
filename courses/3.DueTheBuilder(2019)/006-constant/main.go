package main

import "fmt"

const (
	pi = 3.1415
	f  = 5.0
	f2 = 3.0
)

// all variables are assigned int 1
const (
	a int = 1
	b
	c
	d
)

// iota
const (
	i = iota
	j
	k
	l
)

// mb pow
const (
	m = 1 << 10 * iota
	n
	p
	q
)

// literals
const (
	str      = "string"
	flt      = 3.14
	integrer = 3
	bl       = false
	cmpl     = 3i // (0+3i)
)

func main() {
	fmt.Println("Constants")
	fmt.Println(pi, f, f2)
	// cannot assigned
	//pi = 3.14
	fmt.Println(a, b, c, d) // output: 1 1 1 1
	fmt.Println(i, j, k, l) // output: 0 1 2 3
	fmt.Println(m, n, p, q)

	// study case 1
	var num1, num2 = 5.0, 4
	fmt.Println(num1, num2)
	// fmt.Printf(num1 / num2) // Error: mismatched types float64 and int
	const num3, num4 = 5.0, 4
	fmt.Println(num3 / num4) // No error

	// study case 2
	var x, y, z = 0.1, 0.3, 3.0
	fmt.Println(x + y*z) // 0.999

	const x1, y1, z1 = 0.1, 0.3, 3.0
	fmt.Println(x1 + y1*z1) // 1

	// literals
	fmt.Println(str, flt, integrer, bl, cmpl)

	// other declarations
	var (
		floatPoint  = 1.0e10
		pf          = 1.
		pf2         = .125
		literalRune = 'a'
	)
	fmt.Println(floatPoint, pf, pf2, literalRune)

	// numbers system
	entero := 10
	fmt.Printf("%[1]d %[1]o %[1]x %[1]X %[1]b\n", entero)

	// multiline
	strMultiLine := `Hola, mundo,
Este programa está hecho en Go.`
	fmt.Println(strMultiLine)
}
