package main

import "fmt"

// $ go build -o init2 c.go a.go b.go
func init() {
	fmt.Println("Soy a 1")
}

func init() {
	fmt.Println("Soy a 2")
}

func main() {
	fmt.Println("main()")
}
