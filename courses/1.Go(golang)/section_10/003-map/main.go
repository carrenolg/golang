package main

import "fmt"

func main() {
	m := map[string]int{
		"Messi": 33,
		"Cr7":   35,
	}
	fmt.Println(m)
	// get item from map
	fmt.Println(m["Cr7"])

	// if the key not match
	fmt.Println(m["neymar"]) // you get 0

	// check the key
	value, ok := m["neymar"]
	fmt.Println(value)
	fmt.Println(ok) // bool type

	// checking Idiomatic
	if value, ok := m["neymar"]; ok {
		fmt.Println("The key has a value: ", value)
	} else {
		fmt.Println("The key hasn't a value: ", value)
	}

	// add element to map
	m["Gio"] = 30

	for k, v := range m {
		fmt.Println(k, v)
	}

	// delete one element into the map
	delete(m, "Gio")
	for k, v := range m {
		fmt.Println(k, v)
	}

	// delete key that no exist
	delete(m, "Gio")
}
