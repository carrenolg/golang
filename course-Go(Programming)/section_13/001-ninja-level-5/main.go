package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	// execirse #1
	p1 := person{
		first:   "Anthony",
		last:    "Montropez",
		flavors: []string{`Vanilla`, `Pinaple`},
	}
	p2 := person{
		first:   "Giovanny",
		last:    "Carreño",
		flavors: []string{`Vanilla`, `Pinaple`},
	}

	fmt.Println(p1, p2)

	// anging over the elements
	for _, element := range p1.flavors {
		fmt.Println(element)
	}

	for _, element := range p2.flavors {
		fmt.Println(element)
	}

	// execirse #2
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m[p1.last], m[p2.last])

	for _, element := range m[p1.last].flavors {
		fmt.Println(element)
	}

	for _, element := range m[p2.last].flavors {
		fmt.Println(element)
	}

	// exercise #3
	vh1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Red",
		},
		fourWheel: true,
	}

	vh2 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "White",
		},
		luxury: true,
	}

	fmt.Println(vh1, vh2)

	fmt.Println(vh1.doors, vh2.doors)

	// exersice #4
	// anonymous struct
	vehicle := struct {
		doors int
		color string
	}{
		doors: 2,
		color: "Black",
	}
	fmt.Println(vehicle)
}
