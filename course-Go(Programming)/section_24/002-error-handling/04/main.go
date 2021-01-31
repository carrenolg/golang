package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("../03/names.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
