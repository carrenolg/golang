package main

import "fmt"

func main() {
	// called
	vi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// pass in a function as an argument
	r := even(sum, vi...)
	fmt.Println(r)

}

func sum(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func even(f func(n ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
