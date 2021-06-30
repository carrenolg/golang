package main

import "fmt"

func main() {
	// lab
	fmt.Println("passing arg by value")
	a := 0
	fmt.Println("main:", a)
	foo(a)
	fmt.Println("main:", a)
	fmt.Println("passing arg by reference")
	b := 0
	fmt.Println("main:", b)
	bar(&b)
	fmt.Println("main:", b)
	fmt.Println("passing slice as arg")
	c := []int{1, 2, 3}
	fmt.Println("main:", c)
	ref(c)
	fmt.Println("main:", c)
	fmt.Println("defer")
	fmt.Println(double(5))
}

func foo(a int) {
	a = 10
	fmt.Println("foo:", a)
}

func bar(b *int) {
	*b = 10
	fmt.Println("foo:", *b)
}

func ref(a []int) {
	a[0] = 9
	fmt.Println("ref:", a)
}

func double(x int) (y int) {
	y = x + x
	// defer adds x after the return is executed
	defer func() {
		y += x
	}()
	return y
}
