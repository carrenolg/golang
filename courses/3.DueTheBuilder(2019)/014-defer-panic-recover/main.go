package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// defer
	fmt.Println(">>>>>>>>>>>>>>>>>>>> defer")
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("error: Try open file: %s", err)
		//f.Close()
	}
	defer f.Close() // run before main finished
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
		// f.Close()
	}
	fmt.Println("File content:")
	fmt.Println(string(b))
	// f.Close()

	fmt.Println(">>>>>>>>>>>>>>>>>>>> panic and recover")
	// panic and recover
	// recover gets any panic into a func
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Get panic: %s", r)
		}
	}()

	fileName := ""
	if fileName == "" {
		panic("File name is empty")
	}
	f, err = os.Open(fileName)
	if err != nil {
		log.Fatalf("error: Try open file: %s", err)
		//f.Close()
	}
	defer f.Close() // run before main finished
	b, err = ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
		// f.Close()
	}
	fmt.Println("File content:")
	fmt.Println(string(b))
}
