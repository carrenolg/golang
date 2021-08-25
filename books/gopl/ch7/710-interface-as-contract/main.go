package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

type Leer interface {
	Leer(p []byte) (n int, err error)
}

type Escribir interface {
	Escribir(p []byte) (n int, err error)
}

type LeerEscribir interface {
	Leer
	Escribir
}

func main() {
	fmt.Println("7.1 Interfaces as Contracts")

	// fmt.Printf: writes to the standard output (a file)
	fmt.Printf("it's time %v\n", time.Now())
	// fmt.Sprintf: write to memory buffer
	str := fmt.Sprintf("it's time %v\n", time.Now())
	fmt.Print(str)

	// the two above functions implement the io.Writer interface
	var inter io.Writer
	fmt.Printf("%T\n", inter) // "nil"
	inter = os.Stdout
	fmt.Printf("%T\n", inter) // "*os.File"
	var buf bytes.Buffer
	inter = &buf
	fmt.Printf("%T\n", inter) // "*bytes.Buffer"

	fmt.Println("7.2 Interfaces Types")
	// new interface as combinations of existing ones
	var r Leer
	fmt.Printf("%T\n", r) // "nil"

	fmt.Println("7.2 Interfaces Satisfaction")
}
