package main

import (
	"fmt"
	"log"
	"os"
)

var g = "g"

// func f() {}

var cwd string // global

func init() {
	var err error
	cwd, err := os.Getwd() // local
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}

func main() {
	// scope
	f := "f"
	fmt.Println(f) // "f" local, shadows package-level func f
	fmt.Println(g) // "g" package-level var
	// fmt.Println(h) // compile error: undefined: h

	// explicti and implicit
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println() // \n
	x = "hello!"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
	fmt.Println()
	// scope
	fmt.Println(cwd)
}
