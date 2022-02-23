package main

import "fmt"

const (
	c = "golang"
)

func main() {
	// interpreted literal
	str := "golang"
	fmt.Printf("%T\n", str)
	fmt.Println(str[0], str[len(str)-1])
	fmt.Println(c)
	// raw literal
	raw := `Date: 2012-01-01
Time: 09:15
Day: Monday`
	fmt.Println(raw)
}
