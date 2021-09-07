package main

import (
	"carrenolg.io/books/gopl/ch8/thumbnail"
)

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		/*if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}*/
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// makeThumbnails
	path := "/home/gio10/Pictures/data/"

	names := []string{
		path + "img1.jpg",
		path + "img3.jpg",
	}

	makeThumbnails4(names)

}
