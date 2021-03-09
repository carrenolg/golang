package main

import "fmt"

func main() {
	// Loop, init,condition,post
	// for init; condition;post
	// outner loop and inner loop
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tThe inner loop: %d\n", j)
		}
	}

	// for statement
	// ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
	// 1. For statements with single condition
	z := 10
	for z > 5 {
		fmt.Println("single condition")
		z--
	}
	// 2. For statements with for clause
	for i := 0; i < 10; i++ {
		fmt.Println("for clause")
	}

	for {
		fmt.Println("for infinite")
		break
	}
	// 3. fmt.Println("for clause")

}
