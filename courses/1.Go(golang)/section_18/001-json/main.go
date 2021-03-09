package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	// Go to Json
	p1 := person{
		First: "Luis",
		Last:  "Carreno",
		Age:   60,
	}
	p2 := person{
		First: "Marlene",
		Last:  "Ortiz",
		Age:   51,
	}

	people := []person{p1, p2}

	fmt.Println("Go data:", people)
	fmt.Printf("Go data - type: %T\n", people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Json data:", bs)
	fmt.Printf("Json data - type: %T\n", bs)

	// Json to GO
	js := `[{"First":"Luis","Last":"Carreno","Age":60},{"First":"Marlene","Last":"Ortiz","Age":51}]`

	fmt.Println("Literal string (json):", js)
	fmt.Printf("Literal string (json) - type %T\n", js)

	// convert string to slice of bytes
	by := []byte(js)

	// convert bytes to GO data structure

	var p []person

	err = json.Unmarshal(by, &p)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Go data:", p)
	fmt.Printf("Go data - type %T\n", p)

}
