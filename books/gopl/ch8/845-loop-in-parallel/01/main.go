package main

import (
	"log"

	"carrenolg.io/books/gopl/ch8/thumbnail"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func main() {
	// makeThumbnails
	path := "/home/gio10/Pictures/data/"

	names := []string{
		path + "img1.jpg",
		path + "img3.jpg",
	}

	makeThumbnails(names)
}
