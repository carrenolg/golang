package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
	for _, path := range files {
		data, err := ioutil.ReadFile(path) // returns []byte
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	// prints
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
