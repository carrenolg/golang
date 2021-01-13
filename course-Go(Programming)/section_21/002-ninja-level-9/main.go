package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello, I'm", p.first, p.last)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// implement
	p := person{
		first: "Gio",
		last:  "Carreno",
	}
	saySomething(&p) // we can pass
	// saySomething(p)  // we can not pass

	p.speak()
}
