package main

import (
	"carrenolg.io/books/gopl/ch8/thumbnail"
)

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		/*if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}*/
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

func main() {
	// makeThumbnails
	path := "/home/gio10/Pictures/data/"

	names := []string{
		path + "img1.jpg",
		path + "img3.jpg",
	}

	makeThumbnails3(names)

}
