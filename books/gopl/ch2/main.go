package main

import (
	"fmt"

	"carrenolg.io/books/gopl/ch2/apptemp"
	"carrenolg.io/books/gopl/ch2/popcount"
)

func main() {
	// client
	fmt.Println(apptemp.CtoF(apptemp.BoilingC))
	// package init
	num := popcount.PopCount(uint64(32600))
	fmt.Println(num)
}
