package main

import "fmt"

// $ go build
func init() {
	fmt.Println("Soy a 1")
}

func init() {
	fmt.Println("Soy a 2")
}

func main() {
	fmt.Println("main()")
}
