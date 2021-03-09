package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// define new interface
type human interface {
	speak()
}

// func (r receiver) identifier(parameters) (return(s)) { code }

// implement interface
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "and my LTK's status is:", s.ltk)
}

// implement interface
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am a person")
}

func bar(h human) {
	switch h.(type) {
	case person:
		{
			fmt.Println("I was passed into bar, type: [person]", h.(person).first)
		}
	case secretAgent:
		{
			fmt.Println("I was passed into bar, type: [secretAgent]", h.(secretAgent).first)
		}
	}
}

func main() {
	fmt.Println("Methods")
	// the value p can be of more than one type.
	// in this case: human and secretAgent type
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Gio",
			last:  "Carreño",
		},
		ltk: false,
	}

	p1 := person{
		first: "Anthony",
		last:  "Montropez",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(p1)

	sa1.speak()
	sa2.speak()
	p1.speak()

	// interface method
	bar(sa1)
	bar(sa2)
	bar(p1)
}
