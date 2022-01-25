package main

import (
	"log"

	"carrenolg.io/books/gorest/src/ch7/01-pq/base62"
)

func main() {
	x := 100
	base62String := base62.ToBase62(x)
	log.Println(base62String)
	normalNumber := base62.ToBase10(base62String)
	log.Println(normalNumber)
}
