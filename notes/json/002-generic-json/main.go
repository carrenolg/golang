package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var i interface{}
	fmt.Printf("%T\n", i) // <nil>
	i = "a string"
	fmt.Printf("%T\n", i) // string
	i = 2011
	fmt.Printf("%T\n", i) // int
	i = 2.777
	fmt.Printf("%T\n", i) // float64
	i = "abcd"

	//2. Type assertion
	switch v := i.(type) {
	case int:
		fmt.Println("twice i is:", v*2)
	case float64:
		fmt.Println("the reciplocal of i is:", 1/v)
	case string:
		h := len(v) / 2
		fmt.Println("i swapped by halves is:", v[h:]+v[:h])
	default:
		fmt.Println("i isn't one of the types above")
	}

	//3. Decoding arbitrary JSON data
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	fmt.Printf("%T\n", f) // <nil>
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", f) // map[string]interface{}
	fmt.Println(f)
	m := f.(map[string]interface{})  // get concrte
	fmt.Printf("%T\n", m)            // map[string]interface{}
	fmt.Printf("%T\n", m["Name"])    // string
	fmt.Printf("%T\n", m["Age"])     // float64
	fmt.Printf("%T\n", m["Parents"]) // []interface{}

	// 3. Loop
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is a string", vv)
		case float64:
			fmt.Println(k, "is a float64", vv)
		case []interface{}:
			fmt.Println(k, "is a array", vv)
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
