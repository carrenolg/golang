package main

import "fmt"

func main() {
	// pointers
	// 1. create a pointer
	var p *int
	fmt.Println(p)        // <nil>
	fmt.Printf("%T\n", p) // type [*int]

	// 2. create a variable of type int
	v := 7

	// 3. get pointer of v
	p = &v
	fmt.Println(p)        // 0xc0000160c0
	fmt.Printf("%T\n", p) // type [*int]

	// 4. change the value stored in p
	*p = v * 2
	fmt.Println(p)

	// 5. check value of v
	fmt.Println(v)

	// 6. get value of v through the p
	fmt.Println(*p)

	// 7. create a new pointer
	var p2 *int
	p2 = p
	fmt.Println(p2)

	// 8. check
	fmt.Printf("%T\n", &v)
	fmt.Printf("%v\n", &v)
}
