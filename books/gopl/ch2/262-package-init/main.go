package main

import "fmt"

// package-level variables
var a = b + c
var b = f()
var c = 1

func f() int {
	return c + 1
}

func main() {
	fmt.Println(a, b, c)
}
