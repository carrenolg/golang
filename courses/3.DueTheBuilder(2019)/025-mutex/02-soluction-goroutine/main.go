package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

var path string = "/home/gio10/Documents/books"

var books = map[string]string{
	"alice":   path + "/alice_in_wonderland.txt",
	"iliad":   path + "/the-iliad.txt",
	"quijote": path + "/don-quijote.txt",
}

var mu sync.Mutex
var words = make(map[string]int)

func wordCounter(path string) {
	defer wg.Done()

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
		mu.Lock()
		words[scanner.Text()]++
		mu.Unlock()
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading input:", err)
	}

}

var wg sync.WaitGroup

func main() {
	// mutex
	for i := 0; i < 10; i++ {
		for _, file := range books {
			wg.Add(1)
			go wordCounter(file)
		}
	}

	wg.Wait()
	for w, count := range words {
		fmt.Printf("word: %s, count: %d\n", w, count)
	}
}
