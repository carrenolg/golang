package main

import "fmt"

// h function which returns the product of parameters x and y
func h(x, y int) int {
	return x * y
}

// g function
func g(l, m int) (x int, y int) {
	x = 2 * l
	y = 4 * m
	return
}

func main() {
	fmt.Println(h(g(1, 2)))
}
