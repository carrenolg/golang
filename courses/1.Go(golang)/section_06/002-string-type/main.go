package main

import "fmt"

func main() {
	s := "Hello world!"
	s = `Hello world!`
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	// a string is made up of a slice of bytes.
	// A string value is a sequence of bytes.
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs) // uint8 == byte
	// UTF-8, 1 to 4 byte
	for i := 0; i < len(bs); i++ {
		fmt.Printf("%#U\n", bs[i])
	}
}
