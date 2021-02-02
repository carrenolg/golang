package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// os.Open read a file
	f, err := os.Open("../03/names.txt")
	// f is a pointer to f
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// ioutil.ReadAll reads all file content
	bs, err := ioutil.ReadAll(f)
	// bs = is a byte size
	// []byte = is a byte of string
	if err != nil {
		fmt.Println(err)
	}
	// convert byte to string
	fmt.Println(string(bs))
}
