package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

// implement sort.Interface
func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {
	// 7.6 Sorting with sort.Interface
	var s sort.Interface
	names := []string{"d", "a", "b", "c"}
	fmt.Println(names) // show unordened
	s = StringSlice(names)
	fmt.Printf("type: %T, value: %[1]v\n", s)
	sort.Sort(s)
	fmt.Println(names) // show ordened
}
