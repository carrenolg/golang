package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	input := 3
	for i := 0; i < len(array)-1; i = i + 1 {
		if array[i] == input {
			first := array[0 : i+1]
			second := array[i+1:]
			result := reverse(second)
			fmt.Println(first, result)
		}
	}
}

func reverse(array []int) []int {
	//var newArray []int
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}
