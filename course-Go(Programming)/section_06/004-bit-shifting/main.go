package main

import "fmt"

// using iota with bit shift

func main() {
	// base 10
	kb := 1024
	mb := kb * 1024
	gb := mb * 1024
	fmt.Printf("decimal:%d,\t\tbinary:%b\n", kb, kb)
	fmt.Printf("decimal:%d,\tbinary:%b\n", mb, mb)
	fmt.Printf("decimal:%d,\tbinary:%b\n", gb, gb)
	// using bit shift with iota
	const (
		_    = iota
		kbit = 1 << (iota * 10)
		mbit = 1 << (iota * 10)
		gbit = 1 << (iota * 10)
	)
	fmt.Printf("decimal:%d,\t\tbinary:%b\n", kbit, kbit)
	fmt.Printf("decimal:%d,\tbinary:%b\n", mbit, mbit)
	fmt.Printf("decimal:%d,\tbinary:%b\n", gbit, gbit)
	// types
	fmt.Printf("%T, %v\n", kbit, kbit)
	fmt.Printf("%T, %v\n", mbit, mbit)
	fmt.Printf("%T, %v\n", gbit, gbit)

}
