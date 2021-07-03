package main

import "fmt"

type A struct {
	field1 int
	field2 int
}

type B struct {
	A
	field3 int
	field4 int
}

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader // Promoved Methods
	Writer // Promoved Methods
}

type Document struct {
	nombre string
	File   // Read() and Write() Promoved methods
}

type File struct {
	nombre string
}

func (f File) Read() {
	fmt.Println("Reading File")
}

func (f File) Write() {
	fmt.Println("Writing File")
}

func Print(r Reader) {
	r.Read()
}

func Save(w Writer) {
	w.Write()
}

func PrintSave(ps ReadWriter) {
	fmt.Println("Reading and Writing File")
}

func main() {
	// promoved fields
	v := B{
		A:      A{1, 2},
		field3: 3,
		field4: 4,
	}
	fmt.Printf("%T, %#[1]v\n", v)
	fmt.Println(v.field1, v.field2, v.field3, v.field4)

	// promoved methods
	f := File{nombre: "Data.txt"}
	d := Document{nombre: "Doc", File: File{nombre: "File.txt"}}
	Print(f)
	Print(d)
	Save(f)
	Save(d)
	PrintSave(f)
	PrintSave(d)

	// accessing to promoved fields
	fmt.Println(f.nombre, d.nombre)
	fmt.Println(f.nombre, d.File.nombre)

}
