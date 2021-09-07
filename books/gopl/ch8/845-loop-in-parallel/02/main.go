package main

import (
	"carrenolg.io/books/gopl/ch8/thumbnail"
)

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		/*if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}*/
		go thumbnail.ImageFile(f) // Note: ignoring errors
	}
}

func main() {
	// makeThumbnails
	path := "/home/gio10/Pictures/data/"

	names := []string{
		path + "img1.jpg",
		path + "img3.jpg",
	}

	makeThumbnails2(names)
	// doesn't wait for goroutines to finish
	// program exit
}
