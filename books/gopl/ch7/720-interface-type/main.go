package main

import "fmt"

type Addition interface {
	Add(n int) int
}

type Subtraction interface {
	Sub(n int) int
}

type AddSub interface {
	Addition // embedding interface
	Subtraction
}

// integer: implemets the "AddSub" interface
type integer struct {
	value int
}

func (i integer) Add(n int) int {
	return i.value + n
}

func (i integer) Sub(n int) int {
	return i.value - n
}

func Operation(par AddSub, v int, ope string) {
	switch ope {
	case "a":
		result := par.Add(v)
		fmt.Println(result)
	case "s":
		result := par.Sub(v)
		fmt.Println(result)
	}

}

func main() {
	// 7.2 Interface Types
	num := integer{
		value: 10,
	}
	Operation(num, 5, "s")
	Operation(num, 10, "a")
}
