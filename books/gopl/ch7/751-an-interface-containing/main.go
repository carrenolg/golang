package main

import (
	"bytes"
	"fmt"
	"io"
)

//const debug = true
const debug = false

func f(out io.Writer) {
	// .. do somthing
	fmt.Printf("type: %T, value: %[1]v\n", out)
	if out != nil {
		out.Write([]byte("done!\n"))
		fmt.Print(out)
	}
}

func main() {
	// 7.5.1 Caveat: An Interface Containing a Nil Pointer Is Non-Nil
	var buf *bytes.Buffer
	//var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}
