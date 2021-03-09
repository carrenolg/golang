package main

import "fmt"

func main() {
	// exercise #1
	a := [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("type of array: %T\n", a)

	// exercise #2
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("type of slice: %T\n", s)

	// exercise #3
	fmt.Println(s[1:6])
	fmt.Println(s[3:8])

	// exercise #4
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	// exercise #5
	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	z = append(z[0:3], z[6:]...)
	fmt.Println(z)

	// exercise #6
	states := make([]string, 50, 500)

	/*for i := 0; i < 50; i++ {
		states[i] = fmt.Sprintf("Position %d\n", i)
	}*/

	//fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	states = []string{` Alabama`, ` Alaska`, ` Arizona`,
		` Arkansas`, ` California`, ` Colorado`,
		` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`,
		` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`,
		` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`,
		` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`,
		` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`,
		` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`,
		` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`,
		` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`,
		` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`,
		` West Virginia`, ` Wisconsin`, ` Wyoming`}

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

	// exercise #7
	// multi-dimensional slice
	xxs := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("slice:", i)
		for j, value := range xs {
			fmt.Printf("\t index position: %v \t value: %v\n", j, value)
		}
	}

	// exercise #8
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, slice := range m {
		fmt.Println("map[key]:", k)
		for i, value := range slice {
			fmt.Printf("\t index position: %v \t value: %v\n", i, value)
		}
	}

	// add element
	m[`gio10`] = []string{`futbol`, `Computer Science`, `Science`, `Math`}

	// delete element
	delete(m, `no_dr`)

	fmt.Println(m)

}
