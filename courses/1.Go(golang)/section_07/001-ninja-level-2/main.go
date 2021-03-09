package main

import "fmt"

func main() {
	// exercise #1
	number := 10
	fmt.Printf("decimal:%d,binary:%b,hex:%#x\n", number, number, number)

	// exercise #2
	a := 15 == 25
	b := 'a' <= 'b'
	c := "AAA" != "aaa"
	fmt.Println(a, b, c)

	// exercise #3
	const x = 15
	const y uint8 = 255
	fmt.Println(x, y)

	// exercise #4
	i := 1024
	fmt.Printf("decimal:%d,binary:%b,hex:%#x\n", i, i, i)
	is := i << 1 // bit shift 1 position to the left
	fmt.Printf("decimal:%d,binary:%b,hex:%#x\n", is, is, is)

	// exercise #5
	var s string = `this is a raw string literal`
	fmt.Println(s)

	// exercise #6
	const (
		year1 = iota + 2020
		year2 = iota + 2020
		year3 = iota + 2020
		year4 = iota + 2020
	)
	fmt.Println(year1, year2, year3, year4)

	// exercise #7
	// take the quiz
}
