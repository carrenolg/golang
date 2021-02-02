package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// os.Create creates a new file (in memory)
	// f = pointer to a file
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// close f file
	defer f.Close()

	// pass in "Gio10" as string type
	r := strings.NewReader("Gio10")
	// r = is a *Reader

	// copy r (reader) to f
	io.Copy(f, r)
}
