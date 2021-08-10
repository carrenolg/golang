package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// 3.5.5 Conversions between Strings and Numbers
	x := 123
	// int to string
	y := fmt.Sprintf("str:%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "str:123 123"

	// format numbers, base: 2, 8, 10, 16
	fmt.Println(strconv.FormatInt(int64(x), 2))  // "1111011"
	fmt.Println(strconv.FormatInt(int64(x), 16)) // "7b"

	// string to int (parsing)
	num, err := strconv.Atoi("123")
	if err != nil {
		log.Fatalf("%e\n", err)
	}
	fmt.Println(num) // "123"

	int_64, err := strconv.ParseInt("255", 10, 64)
	if err != nil {
		log.Fatalf("%e\n", err)
	}
	fmt.Println(int_64) // "255"
}
