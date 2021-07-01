package main

import "fmt"

type mensaje string // accesible from external package

type Message struct {
	text string // inaccesible from external package
}

func (t Message) Print() {
	fmt.Println(t.text)
}

func (m mensaje) imprime() {
	fmt.Println(m)
}

func (p *mensaje) cambia(s string) {
	*p = mensaje(s)
}

func (p *Message) change(s string) {
	p.text = s
}

func main() {
	// method
	s := "Hola!"
	fmt.Println(s)
	var m mensaje = "Hola!"
	m.imprime()
	m.cambia("Go!")
	m.imprime()
	// method value (using value)
	v := m.cambia
	v("Method Value in Go!")
	m.imprime()
	// method expression (using type)
	exp := (*mensaje).cambia
	fmt.Printf("%T\n", exp)
	exp(&m, "Method Expression in Go!")
	m.imprime()
	// equivalent
	(*mensaje).cambia(&m, "Method expression equivalent!")
	m.imprime()

	// using struct
	msg := Message{text: "Init value"}
	msg.Print()
	msg.change("Changed value")
	msg.Print()

	pint := new(int) // I do not know
	fmt.Println(pint)
	*pint = 15
	fmt.Println(*pint)
}
