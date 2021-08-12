package main

import (
	"bytes"
	"fmt"

	"carrenolg.io/books/gopl/ch4/appendint"
)

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	// underlying array
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	// create two slices from months
	slice1 := months[4:7]
	slice2 := months[6:9]

	fmt.Println(slice1, slice2) // "[April May June] [June July August]"

	// underlying array changed
	months[6] = "June (Changes)"

	// slices are changed too
	fmt.Println(slice1, slice2) // "[April May June (Changes)] [June (Changes) July August]"

	// slice components: *pointer, len, cap
	fmt.Println("len", len(slice1)) // "3"
	fmt.Println("cap", cap(slice1)) // "9"

	// note: capacity is related to the numbers elements of underlying array

	// slices not comparable between them
	// only with nil
	a := []int{1}
	b := []int{1}
	// fmt.Println(a == b) // compiler error
	fmt.Println(a == nil, b == nil) // "false flase"

	// slices of type []byte can be comparable
	sb1 := []byte{255, 254}
	sb2 := []byte{255, 254}
	eq := bytes.Equal(sb1, sb2)
	fmt.Println(eq) // "true"

	// make: create any slice
	c := make([]int, 5, 10)
	fmt.Println(c, len(c), cap(c)) // "[0 0 0 0 0] 5 10"

	// 4.2.1 The append Function
	fmt.Println("// 4.2.1 The append Function")

	// building a slice of runes
	var runes []rune
	for _, r := range "Hello, Go" {
		runes = append(runes, r)
	}
	fmt.Println(runes)        // "[72 101 108 108 111 44 32 71 111]"
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' 'G' 'o']"

	// other way
	fmt.Printf("%q\n", []rune("Hello, Go")) // "['H' 'e' 'l' 'l' 'o' ',' ' ' 'G' 'o']"

	// copy: built-in function
	underArray := [10]int{9: 10, 5: 10}
	dst := underArray[5:]
	src := underArray[5+1:]
	copy(dst, src)
	// fmt.Println(elements) // "2"
	fmt.Println(dst) // "[0 0]"
	fmt.Println(src) // "[0 0 0 0 0 10]"

	// appendInt
	s := make([]int, 1, 2)
	s[0] = 10
	fmt.Println(s) // "[10]"
	s = appendint.AppendInt(s, 2)
	fmt.Println(s, len(s), cap(s)) // "[10 2] 2 2"

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendint.AppendInt(x, i)
		fmt.Printf("%d cap=%d \t%v\n", i, cap(y), y)
		x = y
	}

	// append
	var n []int
	n = append(n, 1)
	n = append(n, 1, 2)
	n = append(n, n...)
	fmt.Println(n) // "[1 1 2 1 1 2]"

	// 4.2.2 In-Place Slice
	fmt.Println("// 4.2.2 In-Place Slice")

	// slice is change in-place
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // ["one" "three"]
	fmt.Printf("%q\n", data)           // ["one" "three" "three"]

	// slices like stack
	stack := []int{20}
	stack = append(stack, 10) // push
	// get top
	top := stack[len(stack)-1]
	fmt.Println(top) // "10"
	// pop
	stack = stack[:len(stack)-1] // pop 10
	fmt.Println(stack)           // "[20]"
	// remove
	r := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(r, 2)) // "[5 6 8 9]"
}
