package main

import "fmt"

func main() {
	// composite literal
	x := []int{10, 11, 12, 13, 14}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	// index position
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	// slice with range
	// make it easy
	for i, v := range x {
		fmt.Println(i, v)
	}
	// slice with slicing
	for i := 0; i <= cap(x)-1; i++ {
		fmt.Println(x[i])
	}
	// Append and variaric parameters
	y := []int{200, 300, 400, 500}
	x = append(x, y...)
	fmt.Println(x)
}
