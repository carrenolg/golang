package main

import "fmt"

func main() {
	fmt.Println("Defer: ")
	defer alwaysToEnd()
	foobar()
}

func alwaysToEnd() {
	fmt.Println("the end, using defer keyword")
}

func foobar() {
	fmt.Println("Hello from foobar")
}
