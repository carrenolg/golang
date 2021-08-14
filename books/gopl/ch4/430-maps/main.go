package main

import (
	"fmt"
	"sort"
)

func main() {
	// def: maps is a reference a hash table
	// map operations
	// create
	ages := make(map[string]int)
	fmt.Println(ages) // "map[]"

	// map literal
	ages = map[string]int{
		"jennifer": 23,
		"giovanny": 31,
	}
	fmt.Println(ages) // "map[giovanny:31 jennifer:23]"

	// empty map
	var m map[string]int
	fmt.Println(m) // "map[]"

	// add element
	ages["danna"] = 21
	fmt.Println(ages["danna"]) // "21"

	// get element
	value := ages["giovanny"]
	fmt.Println(value) // "31"

	// delete element
	delete(ages, "danna")
	fmt.Println(ages["danna"]) // "0"

	// safe operations, return zero value
	// "bob" is not a key
	value = ages["bob"]
	fmt.Println(value) // "0"

	// a map element is not a variable
	// _ = &ages["giovanny"] // compile error

	// enumerate key/value pairs
	// unordered
	for _, v := range ages {
		fmt.Println(v)
	}

	// ordered by name (explicitly or manually)
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s:\t%d\n", name, ages[name])
		// fmt.Println(name)
	}

	// map not init (nil)
	var x map[string]int
	fmt.Println(x == nil)    // "true"
	fmt.Println(len(x) == 0) // "true"
	// x["gio"] = 10 // panic: nil map

	// check if a key exist
	age, ok := ages["jennifer"]
	if ok {
		fmt.Println(ok, age) // "true 23"
	} else {
		fmt.Println(ok, age)
	}

	age, ok = ages["bob"]
	if ok {
		fmt.Println(ok, age)
	} else {
		fmt.Println(ok, age) // "flase 0"
	}
}
