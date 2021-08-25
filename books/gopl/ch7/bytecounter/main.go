package main

import "fmt"

type ByteCounter int

// ByteCounter satisfies the io.Write
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {
	// bytecounter
	var c ByteCounter
	(&c).Write([]byte("hello")) // call: (*ByteCounter).Write
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
