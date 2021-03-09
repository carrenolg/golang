package main

import "fmt"

func main() {
	// Composite Literal
	// a := type{values}
	x := []int{10, 11, 12, 13, 14}
	fmt.Println(x)
	// len of the slice
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// index position
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])

	// slice with range
	// make it easy
	// show both the index and value
	for i, v := range x {
		fmt.Println(i, v)
	}
	// slicing a slice
	fmt.Println(x[1:])
	fmt.Println(x[0:1]) // get only the first element
	fmt.Println(x[2:2]) // get a Null/empty slice

	for i := 0; i <= cap(x)-1; i++ {
		fmt.Println(x[i])
	}
	// Append and variaric parameters
	y := []int{200, 300, 400, 500}
	z := []int{10, 11, 12}
	x = append(x, y...)
	// x = append(x, z) Error
	fmt.Println(x)
	fmt.Println(z)

	// deling from a slice
	// delete only one element
	// this case the element "300" with index = 6
	x = append(x[:6], x[7:]...)
	fmt.Println(x)

	// make
	var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints
	fmt.Println(v)

	m := make([]int, 10, 100)
	fmt.Println(m)

	n := make([]string, 10, 100)
	fmt.Println(n)

	b := make([]bool, 10, 100)
	fmt.Println(b)

	// len and capacity
	fmt.Println(len(v))
	fmt.Println(cap(v))

	fmt.Println(len(b))
	fmt.Println(cap(b))

	// testing capacity
	t := make([]int, 2, 5)
	fmt.Println(t)

	t[0] = 10
	t[1] = 20
	fmt.Println(t)

	//t[2] = 30 // index out of range
	t = append(t, 30)
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(cap(t))

	// duplicate capacity
	t = append(t, 40, 50, 60, 70)
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(cap(t))

	// multi-dimensional slice
	dc := []string{"Danna", "Valentina", "Age: 20"}
	fmt.Println(dc)
	lc := []string{"Luis", "Giovanny", "Age: 30"}
	fmt.Println(lc)
	multiSlice := [][]string{dc, lc}
	fmt.Println(multiSlice)
}
