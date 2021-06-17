package main

import (
	"fmt"
)

func main() {
	// slices declaration
	var s []int // nil zero value
	imprime(s, "s")
	// literal empty, no nil
	var q = []int{}
	imprime(q, "q")
	// slice literal
	var p = []int{1, 2, 3}
	imprime(p, "p")

	// create new slice from slice
	var a = []int{0, 1, 2, 3, 4, 5}
	imprime(a, "a")
	b := a[1:5] // get index 1 to 4
	imprime(b, "b")
	c := a[:] // get all index
	imprime(c, "c")

	// childs and parents sharing same array base
	var parent = []int{10, 20, 30, 40, 50}
	imprime(parent, "parent")
	child := parent[0:2] // get first two elements
	imprime(child, "child")
	parent[1] = 200000
	// if change the parent, then the child too
	imprime(parent, "parent")
	imprime(child, "child")

	// capacity of slice
	var w = []int{0, 1, 2, 3, 4, 5}
	w1 := w[1:4]
	imprime(w, "w")
	imprime(w1, "w1")
	// we extende capacity by 1
	w1 = w1[:len(w1)+1]
	w1 = w1[:cap(w1)]
	imprime(w1, "w1")

	// append func
	w = append(w, 11, 12, 13, 14)
	imprime(w, "w")

	// make
	var m = []int{0, 1, 2, 3, 4, 5}
	l := make([]int, 6)
	copy(l, m) // l = m
	imprime(m, "m")
	imprime(l, "l")
	m[2] = 100
	imprime(m, "m")
	imprime(l, "l")
	// m and l have different memory location
	// fmt.Println(m == l) # invalido: slice can only be compared to nil

	fmt.Println("\nLab #######")

	// reference to an underlying array
	// slice has its own reference, this isn't same to underlying array
	var r = []int{15, 25, 35, 45, 55}
	f := r[1:3]
	fmt.Printf("ref(r): %p\n", &r) // slice self-reference
	fmt.Printf("ref(f): %p\n", &f) // slice self-reference
	r[2] = 300
	imprime(r, "r")
	imprime(f, "f")

}

func imprime(s []int, n string) {
	fmt.Println("******************************")
	fmt.Printf("slice %s, len: %d, cap: %d, nil: %t\n", n, len(s), cap(s), s == nil)
	fmt.Print("{")
	for i, v := range s {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%d: %v", i, v)
	}
	fmt.Println("}")
}
