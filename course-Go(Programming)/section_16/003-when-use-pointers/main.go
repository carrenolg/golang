package main

import "fmt"

func main() {
	a := 15
	fmt.Println("a before:", a)
	foo(a)
	fmt.Println("a after:", a)

	a = 15
	fmt.Println("a before:", a)
	bar(&a)
	fmt.Println("a after:", a)
}

func foo(b int) {
	b = b * b
	fmt.Println(b)
}

func bar(b *int) {
	*b = *b * *b // dereference
	fmt.Println(*b)
}
