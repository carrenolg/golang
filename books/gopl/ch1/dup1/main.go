package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Finding duplicate lines in a file
	counts := make(map[string]int)
	file, err := os.Open("alice_in_wonderland.txt")
	if err != nil {
		log.Println("error al abrir archivo", err)
	}
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
