package main

import "fmt"

func main() {
	// if
	x := 20
	// no idiomatic
	if x == 10 {
		fmt.Println("x == 10")
	} else {
		fmt.Println("x != 10")
	}
	// idiomatic
	if x == 10 {
		fmt.Println("x == 10")
		return // exit from main
	}
	fmt.Println("x != 10")

	// call
	if r, err := f(5); err == nil {
		fmt.Println(r)
		return
	} else {
		fmt.Println(err)
	}

}

func f(n int) (int, error) {
	// idiomatic
	if n == 10 {
		return n, nil
	}
	return 0, fmt.Errorf("error %d operation", 010)
}
