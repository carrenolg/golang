package main

import "fmt"

func main() {
	// map
	// declaration
	var m map[string]int // nil
	imprime(m, "m")
	fmt.Println(m["Juan"]) // 0

	// map literals
	var p = map[string]int{
		"Juan":  20,
		"Pedro": 50,
	}
	imprime(p, "p")

	// using make
	var q = make(map[string]int) // no nil
	imprime(q, "q")
	fmt.Println(q["Pedro"]) // 0

	// get item
	var i = map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	}
	fmt.Println(i["a"], i["b"], i["c"], i["d"])

	// add
	i["z"] = 27

	// delete item
	delete(i, "c")
	fmt.Println(i["a"], i["b"], i["c"], i["d"])

	// check item exist
	e, ok := i["c"]
	fmt.Println(e, ok) // false

	// behaivor map nil or no init
	var w map[string]int
	fmt.Println(w["a"], len(w)) // 0 0
	//>> w["b"] = 15                 // runtime error

	// notes
	//1. if item no exist return the zero value of item
	//2. map's keys are not sorted
	//3. make is using for map initialitation

}

func imprime(m map[string]int, n string) {
	fmt.Println("******************************")
	fmt.Printf("map %s, len: %d, nil: %t\n", n, len(m), m == nil)
	fmt.Print("{")
	for k, v := range m {
		fmt.Printf("%s: %v, ", k, v)
	}
	fmt.Println("}")
}
