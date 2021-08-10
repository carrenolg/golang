package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// 3.6 Constants
	const pi = 3.14159
	fmt.Println(pi)
	// multi declaration
	const (
		e = 2.718281
		g = 9.8
	)
	fmt.Println(e, g)

	// run time vs compile time
	const d = 10
	// const r = d / 0 // Dive by 0
	var p [d]byte
	fmt.Println(p) // "[0 0 0 0 0 0 0 0 0 0]"

	// explicit type and implicit type
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	// generator iota
	const (
		Bit1 = 1 << iota
		Bit2
		Bit3
		Bit4
		Bit5
		Bit6
		Bit7
		Bit8
	)
	fmt.Printf("%08b %08b\n", Bit1, Bit8)

	// 3.6.2 Untyped Constants
	fmt.Println("// 3.6.2 Untyped Constants")

	const (
		Base = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
		PiB
		EiB
		ZiB
		YiB
	)
	fmt.Println(Base, YiB/ZiB) // "1 1024"

	// var or const

	const Pi64 float64 = math.Pi
	// var x float32 = Pi64, required conversion
	var x float32 = float32(Pi64)
	fmt.Println(x, Pi64)

	// operations
	var f float64 = 212
	fmt.Println((f - 32))

	// study case
	// fmt.Println(5 / 9 * (f - 32)) // "0"
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float
	fmt.Println(5.0 / 9.0 * 180)      // "0.5555555555555556"
}
