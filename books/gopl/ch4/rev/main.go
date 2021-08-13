package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func revPtr(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// rev
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"

	// rev with array pointer
	b := [...]int{0, 1, 2, 3, 4, 5}
	revPtr(&b)
	fmt.Println(b) // "[5 4 3 2 1 0]"

	// rotate
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	fmt.Println(s) // "[1 0 2 3 4 5]"
	reverse(s[2:])
	fmt.Println(s) // "[1 0 5 4 3 2]"
	reverse(s[:])
	fmt.Println(s) // "[2 3 4 5 0 1]"

	// rotate simple pass (Exercise 4.4)
	r := []int{0, 1, 2, 3, 4, 5}
	r2 := make([]int, len(r))
	n, m := copy(r2[0:4], r[2:]), copy(r2[4:], r[0:2])
	fmt.Println(n, m)
	fmt.Println(r)
	fmt.Println(r2)
}
