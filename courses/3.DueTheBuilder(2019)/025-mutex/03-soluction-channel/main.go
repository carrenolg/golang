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

//var mu sync.Mutex
var words = make(map[string]int)

func wordCounter(path string, out chan<- map[string]int) {
	defer wg.Done()

	words := make(map[string]int)

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
		//mu.Lock()
		words[scanner.Text()]++
		//mu.Unlock()
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading input:", err)
	}
	// send to channel
	out <- words
}

var wg sync.WaitGroup

func main() {
	// channel
	output := make(chan map[string]int)
	for i := 0; i < 1; i++ {
		for _, file := range books {
			wg.Add(1)
			go wordCounter(file, output)
		}
	}

	//wg.Wait()
	go func() {
		wg.Wait()
		close(output)
	}()

	// receiver from channel
	for m := range output {
		for p, c := range m {
			words[p] += c
		}
	}

	// print result
	for w, count := range words {
		fmt.Printf("word: %s, count: %d\n", w, count)
	}
}
