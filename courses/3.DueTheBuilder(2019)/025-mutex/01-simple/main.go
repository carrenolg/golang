package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var path string = "/home/gio10/Documents/books"

var books = map[string]string{
	"alice":   path + "/alice_in_wonderland.txt",
	"iliad":   path + "/the-iliad.txt",
	"quijote": path + "/don-quijote.txt",
}

var words = make(map[string]int)

func wordCounter(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	// read buffer
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading input:", err)
	}

}

func main() {
	// mutex
	for i := 0; i < 1000; i++ {
		for _, file := range books {
			wordCounter(file)
		}
	}

	for w, count := range words {
		fmt.Printf("word: %s, count: %d\n", w, count)
	}
}
