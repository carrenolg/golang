package main

import "fmt"

func main() {
	fmt.Println("Recursion")
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
