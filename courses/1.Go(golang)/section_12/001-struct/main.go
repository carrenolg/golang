package main

import "fmt"

// define a new struct
type person struct {
	first string
	last  string
	age   int
}

// embedded struct
type secretAgent struct {
	person //Anonymous field
	ltk    bool
}

func main() {
	// create a struct of type person
	p1 := person{
		first: "Luis",
		last:  "Carreño",
		age:   30,
	}
	// print out
	fmt.Println(p1)
	// print fields out
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Printf("%T\n", p1)

	// embedded struct
	sa1 := secretAgent{
		person: person{
			first: "Luis",
			last:  "Carreño",
			age:   30,
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)                      // fields promoved
	fmt.Println(sa1.person.first, sa1.person.last, sa1.person.age, sa1.ltk) // name space
	fmt.Printf("%T\n", sa1)

	// anonymous struct
	as := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(as)
	fmt.Printf("%T\n", as)
	fmt.Println(as.first, as.last, as.age)

}

/*
fmt.Println(7 * 6 * 5 * 4 * 3 * 2 * 1)
	n := factorial(7)
	fmt.Println(n)
func factorial(n int) int {
	var result int = 1
	for ; n > 0; n-- {
		result *= n
	}
	return result
}
*/
