package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// main
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	fmt.Println(text)
}
