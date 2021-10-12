package main

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// routes
	r := httprouter.New()
	fmt.Println(r)
}
