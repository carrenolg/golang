package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"carrenolg.io/books/gopl/ch8/thumbnail"
)

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func sendNames(out chan<- string, names []string) {
	for _, v := range names {
		out <- v
	}
	close(out)
}

func main() {
	// makeThumbnails
	path := "/home/gio10/Pictures/data/"
	ch := make(chan string)
	names := []string{
		path + "img1.jpg",
		path + "img3.jpg",
	}

	go sendNames(ch, names)
	total := makeThumbnails6(ch)
	fmt.Println("Total:", total)
}
