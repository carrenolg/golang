package main

import (
	"fmt"
	"log"

	"carrenolg.io/books/gopl/ch8/thumbnail"
)

type item struct {
	thumbfile string
	err       error
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {

	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		/*if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}*/
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func main() {
	// makeThumbnails
	path := "/home/gio10/Pictures/data/"

	names := []string{
		path + "img1.jpg",
		path + "img3.jpg",
	}

	files, err := makeThumbnails5(names)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(files)

}
