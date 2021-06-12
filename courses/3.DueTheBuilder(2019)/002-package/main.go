package main

import (
	"fmt"

	"carrenolg.io/mate" // local package ~/go/src/
)

func main() {
	fmt.Println("Working with Go's packages")
	r := mate.Suma(5, 5)
	fmt.Println(r)
}
