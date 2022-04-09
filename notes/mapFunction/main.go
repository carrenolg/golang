package main

import "fmt"

//type: func(int) int

//mapFunc: recives a function as parameter
// and returns a value of type function
func mapFunc(fn func(int) int) func(int) int {
	body := func(num int) int {
		return fn(num)
	}
	return body
}

func main() {
	// declare a value of type: func(int) int
	doble := func(x int) int {
		return x * x
	}
	f := mapFunc(doble)
	fmt.Println(f(5))
}

//output: 25
