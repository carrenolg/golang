package main

import (
	"fmt"
	"os"
)

func main() {
	// 7.11 Discriminating Errors with Type Assertions
	_, err := os.Open("/no/such/file")
	fmt.Println(err)
	fmt.Printf("%#v\n", err)
	// type: *os.PathError

	// check what type error is
	_, err = os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err)) // "true"
}
