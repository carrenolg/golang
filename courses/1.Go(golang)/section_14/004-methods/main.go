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

// func (r receiver) identifier(parameters) (return(s)) { code }

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "and my LTK's status is:", s.ltk)
}

func main() {
	fmt.Println("Methods")
	p := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p2 := secretAgent{
		person: person{
			first: "Gio",
			last:  "Carreño",
		},
		ltk: false,
	}

	fmt.Println(p)
	fmt.Println(p2)

	p.speak()
	p2.speak()

}
