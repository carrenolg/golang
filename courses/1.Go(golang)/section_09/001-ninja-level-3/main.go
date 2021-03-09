package main

import "fmt"

func main() {
	// Ninja level 3
	// exercise # 1
	for n := 1; n <= 100; n++ {
		fmt.Println(n)
	}
	// exercise # 2
	// Latin Alphabet:Uppercase
	for code := 65; code <= 90; code++ {
		fmt.Println(code)
		for range []int{0, 1, 2} {
			fmt.Printf("\t%#U\n", code)
		}
	}
	// exercise # 5
	for n := 10; n <= 100; n++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", n, n%4)
	}
	// exercise #

}
