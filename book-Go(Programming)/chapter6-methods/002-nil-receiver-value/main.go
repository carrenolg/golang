package main

import "fmt"

type intList struct {
	value int
	tail  *intList
}

func (list *intList) sum() int {
	if list == nil {
		return 0
	}
	return list.value + list.tail.sum()
}

func main() {
	l := intList{}
	fmt.Println(l)

}
