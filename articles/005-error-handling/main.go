package main

import "fmt"

func CheckError(err error) bool {
	if err != nil {
		return true
	} else {
		return false
	}
}

func div(dividen, divisor int) (int, error) {
	if divisor < 0 {
		return -1, fmt.Errorf("divisor with value %d is smaller that zero", divisor)
	} else {
		return dividen / divisor, nil
	}
}

func main() {
	// simple error handling
	value, err := div(5, -2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Result:", value)

	// advanced error handling
	if value, err := div(5, -2); CheckError(err) {
		fmt.Println(err)
	} else {
		fmt.Println("Result:", value)
	}
}
