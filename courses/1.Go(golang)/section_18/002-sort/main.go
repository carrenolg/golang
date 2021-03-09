package main

import (
	"fmt"
	"sort"
)

// Person is a new type
type Person struct {
	Name string
	Age  int
}

// ByAge implement interface
type ByAge []Person

// Len method
func (a ByAge) Len() int {
	return len(a)
}

// Swap method
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less method
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	// sort package
	// int
	a := []int{0, 3, 4, 19, 2, 3, 4, 200, 6}
	sort.Ints(a)
	fmt.Println(a)

	// string
	b := []string{"a", "z", "c", "g", "t", "l", "p"}
	sort.Strings(b)
	fmt.Println(b)

	// sort custom (struct)
	p1 := Person{Name: "Gio", Age: 30}
	p2 := Person{Name: "Antho", Age: 30}
	p3 := Person{Name: "Angel", Age: 27}
	p4 := Person{Name: "Emerson", Age: 26}
	p5 := Person{Name: "ArioBula", Age: 31}

	people := []Person{p1, p2, p3, p4, p5}

	fmt.Println(people)

	sort.Sort(ByAge(people))

	fmt.Println(people)

}
