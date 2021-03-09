package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type square struct {
	long  float32
	width float32
}

type circle struct {
	radio float32
}

type shape interface {
	area() float32
}

func main() {
	// exercise #3
	defer deferFunc()

	// exercise #2
	n := []int{10, 11, 12}
	sum := foo(n...)
	fmt.Println(sum)
	sum2 := bar(n)
	fmt.Println(sum2)

	// exercise #3
	p := person{
		first: "Gio",
		last:  "Carreño",
		age:   30,
	}
	p.speak()

	// exersice #4
	s := square{
		long:  20,
		width: 10,
	}
	c := circle{
		radio: 2.5,
	}
	// method interface
	info(s)
	info(c)

}

func foo(vp ...int) int {
	total := 0
	for _, v := range vp {
		total += v
	}
	return total
}

func bar(p []int) int {
	total := 0
	for _, v := range p {
		total += v
	}
	return total
}

func deferFunc() {
	fmt.Println("Hello from deferFunc")
}

func (p person) speak() {
	fmt.Println("Hello, I'm", p.first, p.last, "age:", p.age)
}

// implement interface
func (s square) area() float32 {
	return s.long * s.width
}

// implement interface
func (c circle) area() float32 {
	return 3.14159 * (c.radio * c.radio)
}

func info(s shape) {
	fmt.Println("Area:", s.area())
}

/*func info(s shape) {
	switch s.(type) {
	case square:
		{
			fmt.Println("Area of square:", s.(square).area())
		}
	case circle:
		{
			fmt.Println("Area of circle:", s.(circle).area())
		}
	}
} */
