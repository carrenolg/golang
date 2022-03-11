package main

import (
	"log"

	"carrenolg.io/books/godp/ch5/strategy"
)

//var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

func main() {
	// flag.Parse()

	/*var activeStrategy strategy.OutputStrategy

	switch *output {
	case "console":
		activeStrategy = &strategy.TextSquare{}
	case "image":
		activeStrategy = &strategy.ImageSquare{DestinationFilePath: "/tmp/image.jpg"}
	default:
		activeStrategy = &strategy.TextSquare{}
	}

	err := activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}*/
	var text strategy.TextSquare
	OutPut(text)
	var image strategy.ImageSquare
	OutPut(image)
}

func OutPut(activeStrategy strategy.OutputStrategy) {
	switch activeStrategy.(type) {
	case strategy.TextSquare:
		activeStrategy = &strategy.TextSquare{}
	case strategy.ImageSquare:
		activeStrategy = &strategy.ImageSquare{DestinationFilePath: "/tmp/image.jpg"}
	default:
		activeStrategy = &strategy.TextSquare{}
	}

	err := activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
