package main

import "fmt"

func main() {
	// playing with functions
	foo()

	// Everything in GO is passed by VALUE
	bar("Oliver")

	// return simple value
	r := sum(1000, 1)
	fmt.Println(r)

	// return multiple values
	a, b := mouse("Luis", "Carreño")
	fmt.Println(a)
	fmt.Println(b)

}

func foo() {
	fmt.Println("Hello from the [foo] function")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func sum(a, b int) int {
	return a + b
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ` `, ln, `, says "hello"`)
	b := true
	return a, b
}
