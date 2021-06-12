package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// string: secuence of bytes
	// Go codes string by unicode
	// unicode uses between 1 and 4 bytes to encode characters
	// utf8 is a variable len shema
	s := "Hola mundo!"
	fmt.Println()
	fmt.Printf("s tiene %d bytes \n", len(s))
	fmt.Printf("s tiene %d runes \n", utf8.RuneCountInString(s))

	s = "Go en español"
	fmt.Println()
	fmt.Printf("s tiene %d bytes \n", len(s))
	fmt.Printf("s tiene %d runes \n", utf8.RuneCountInString(s))

	encabezado()
	for i := 0; i < len(s); i++ {
		fila(i, s[i])
	}

	encabezado()
	for i, v := range s {
		fila(i, v)
	}
}

func encabezado() {
	fmt.Println()
	fmt.Printf("%8s\t%8s\t%8s\t%8s\n", "índice", "código", "letra", "tipo")
	fmt.Printf("%8[1]s\t%8[1]s\t%8[1]s\t%8[1]s\n", "--------")
}

func fila(i int, x interface{}) {
	fmt.Printf("%8[1]d\t%8[2]d\t%8[2]c\t%8[2]T\n", i, x)
}
