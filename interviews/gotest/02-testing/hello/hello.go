package main

import (
	"fmt"

	"carrenolg.io/interviews/hello/morestring"
	"github.com/google/go-cmp/cmp"
)

func main() {
	//fmt.Println("Hello, world.")
	fmt.Println(morestring.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
