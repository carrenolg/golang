package main

import "fmt"

type book struct {
	title string
	price float64
}

// error
func (r book) Error() string {
	return "Error: this is a error"
}

func main() {
	// under-standing
	b := book{
		title: "Golang",
		price: 12.5,
	}
	fmt.Println(b.title)
	fmt.Println(b.price)
	fmt.Printf("%T\n", b)
	// conversion
	e := error(b)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
}
