package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 3 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err != nil {
			return nil // sucess
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	// 5.4 Erros
	// additional result
	var m map[string]int
	value, ok := m["k"]
	if !ok {
		fmt.Fprintf(os.Stderr, "keys 'k' does no exist %v\n", value)
	}

	// retrying
	if err := WaitForServer("http://golang.org"); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
	}

	// exit program
	if err := WaitForServer("http://golang.org"); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}

}
