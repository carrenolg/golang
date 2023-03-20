package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, w, _ := os.Pipe()
	w.Write([]byte("hola"))
	_ = w.Close()
	out, _ := io.ReadAll(r)
	fmt.Println(string(out))
}
