package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// read file
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("error with file")
	}
	defer f.Close()

	// get content
	// bs, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
	}
	// convert byte to string
	//  fmt.Println(string(bs))

	// read file by line
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for k, v := range strings.Split(text[0], ",") {
		fmt.Println(k, v)
	}

}
