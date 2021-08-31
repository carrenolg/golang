package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// 7.10 Type Assertions
	// w: interface value
	var w io.Writer = nil
	// get dynamic type and value
	w = os.Stdout
	fmt.Printf("type: %T, value: %[1]v\n", w)
	// operation
	f := w.(*os.File)
	// f: interface value
	fmt.Printf("type: %T, value: %[1]v\n", f)
	// c := w.(*bytes.Buffer)
	// panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
	// dinamic type are not equal

	// abstract type
	w = os.Stdout // io.Writer
	rw := w.(io.ReadWriter)
	fmt.Printf("type: %T, value: %[1]v\n", rw)
	fmt.Println(w == rw)
	w.Write([]byte("hello\n"))
	rw.Write([]byte("hello\n"))
	rw.Read([]byte(""))

	// check dynamic type of an interface value
	var i io.Writer = os.Stdout
	f, ok := i.(*os.File)
	fmt.Println(f, ok) // "true"
	b, ok := i.(*bytes.Buffer)
	fmt.Println(b, ok) // "false"

}
