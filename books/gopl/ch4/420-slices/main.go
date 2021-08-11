package main

import (
	"bytes"
	"fmt"
)

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
}
