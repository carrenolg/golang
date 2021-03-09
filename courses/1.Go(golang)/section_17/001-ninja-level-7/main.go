package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	// exercise #1
	v := "golang"
	fmt.Printf("%T\n", &v)
	p := &v
	fmt.Println(p)

	// exercise #2
	b := person{
		first: "Anthony",
		last:  "Montropez",
		age:   30,
	}
	changeMe(&b)
	fmt.Println(b.first, b.last, b.age)

}

func changeMe(p *person) {
	fmt.Println(p.first, p.last, p.age)
	// change data
	// use dereference
	p.first = "Gio"
	p.last = "Carreño"
	p.age = 30
	// other way
	(*p).first = "Gio"
	(*p).last = "Carreño"
	(*p).age = 30

}
