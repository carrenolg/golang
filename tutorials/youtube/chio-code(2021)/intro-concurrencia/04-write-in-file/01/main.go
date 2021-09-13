package main

import (
	"fmt"
	"os"
)

func main() {
	// write in file
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	final := 16777215

	for i := 0; i < final; i++ {
		_, err := f.WriteString(fmt.Sprintf("%06x\n", i))
		if err != nil {
			panic(err)
		}
	}
}
