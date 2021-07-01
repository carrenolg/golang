package main

import "fmt"

func main() {
	// interface
	d := Dog{}
	c := Cat{}
	d.InfoDog()
	c.InfoCat()
	fmt.Println(">>>>>>>>>> using interface")
	Info(d)
	Info(c)
	fmt.Println(">>>>>>>>>> empty interface")
	// empty interface
	Empty(d)
	Empty(c)
}

type Dog struct{}
type Cat struct{}

// interface
type Animal interface {
	Sound()
}

func (d Dog) Sound() {
	fmt.Println("Woof Woof")
}

func (d Cat) Sound() {
	fmt.Println("Meow Meow")
}

func (d Dog) InfoDog() {
	fmt.Println("El animal suena:")
	d.Sound()
}

func (c Cat) InfoCat() {
	fmt.Println("El animal suena")
	c.Sound()
}

// fun for interface
func Info(a Animal) {
	fmt.Println("El animal suena")
	a.Sound()
}

// interface empty
func Empty(s interface{}) {
	fmt.Printf("Type %T implements the interface empty\n", s)
}
