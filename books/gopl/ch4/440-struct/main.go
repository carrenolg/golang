package main

import (
	"fmt"
	"time"
)

func main() {
	// 4.4 Struct
	// def: a struct is an aggregate data type
	// declare
	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time
		Position  string
		Salary    int
		ManagerID int
	}
	// int
	var dilbert Employee
	dilbert.Salary -= 500
	fmt.Println(dilbert) // "{0   0001-01-01 00:00:00 +0000 UTC  -500 0}"

	// accesing by pointer
	position := &dilbert.Position
	*position = "Senior Dev Go" + *position
	fmt.Println(dilbert.Position) // "Senior Dev Go"

	// empty struct, no fields
	seen := make(map[string]struct{})
	if _, ok := seen["s"]; !ok {
		seen["s"] = struct{}{}
	}
	fmt.Println(seen) // "map[s:{}]"

	// 4.4.1 Struct Literals

}
