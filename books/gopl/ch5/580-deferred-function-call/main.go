package main

import "fmt"

func doble(x int) (result int) {
	defer func() {
		fmt.Printf("doble(%d) = %d\n", x, result)
	}()
	return x + x
}

func main() {
	// 5.8 Deferred Function Calls
	// run after return statements
	r := doble(4)  // "doble(4) = 8"
	fmt.Println(r) // "8"

}
