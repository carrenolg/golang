package main

import "fmt"

func main() {
	// For statement
	// 1. SingleCondition
	fmt.Println("1. SingleCondition")
	a := 0
	for b := 10; a < b; {
		fmt.Println(a)
		a++
	}

	// 2. ForClause
	// init: assignment
	// post: increment or decrement
	// condition: SingleCondition
	fmt.Println("2. ForClause")
	for ; a < 20; a++ {
		fmt.Println(a)
	}

	// 3. RangeClause
	fmt.Println("3. RangeClause")
	r := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for a = range r {
		fmt.Println(a)
	}

	// range clause without the iteration variables
	for range r {
		fmt.Println("iteration")
	}

	// examples
	fmt.Println("example 1")
	var testdata *struct {
		a *[7]int
	}
	for i := range testdata.a {
		// testdata.a is never evaluated; len(testdata.a) is constant
		// i ranges from 0 to 6
		fmt.Println(i)
	}
	fmt.Println("example 2")
	var x [10]string
	for i, s := range x {
		// type of i is int
		// type of s is string
		// s == a[i]
		fmt.Printf("%d %s\n", i, s)
	}

	fmt.Println("example 3")
	var key string
	var val interface{} // element type of m is assignable to val
	m := map[string]int{"mon": 0, "tue": 1, "wed": 2, "thu": 3, "fri": 4, "sat": 5, "sun": 6}
	for key, val = range m {
		fmt.Printf("%v %v\n", key, val)
	}

}
