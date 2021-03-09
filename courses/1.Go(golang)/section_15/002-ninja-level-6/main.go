package main

import "fmt"

func main() {
	// exercise #6
	f := func(pow int) int {
		return pow * pow
	}
	fmt.Println("Pow:", f(5))
	fmt.Println("Pow:", f(f(5)))
	fmt.Println("Pow:", f(f(f(5))))
	// other anony func
	func() {
		fmt.Println("This is a anony func")
	}()

	// exercise #8
	vp := []int{10, 20, 30}
	// return func
	fu := sum(vp...)
	// run func
	total := fu()
	// print returned value out
	fmt.Println(total)
}

func sum(n ...int) func() int {
	return func() int {
		total := 0
		for _, v := range n {
			total += v
		}
		return total
	}
}
