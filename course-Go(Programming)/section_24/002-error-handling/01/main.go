package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello")
	// n = numbers of bytes
	// err = type error or nil
	if err != nil {
		// nil = zero or empty
		fmt.Println(err)
	}
	fmt.Println(n)
}
