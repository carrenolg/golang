package main

import (
	"errors"
	"fmt"
)

func main() {
	// functions
	r := multi(5, 5)
	fmt.Println("r:", r)

	// multi return
	change, percentChange := getStockPriceChange(75000.0, 100000.0)
	fmt.Println("multi return:", change, percentChange)

	// get error
	change, percentChange, err := getStockPriceChangeWithError(0, 0)
	fmt.Println("get error:", change, percentChange, err)

	// named return values
	change, percentChange = getStockPriceChangeWithNamedReturnValues(75000.0, 100000.0)
	fmt.Println("named return values:", change, percentChange)

	// ignore parameter
	r = sum(5, 5, 5)
	fmt.Println("ignore par:", r)

	// function value
	f := sum
	fmt.Println("func value:", f(5, 5, 5))
	fmt.Printf("type: %T\n", f) // getting func signatute

	// analysis the fun signature
	s := show
	fmt.Printf("type: %T\n", s)
	p := product
	fmt.Printf("type: %T\n", p)
	n := nothing
	fmt.Printf("type: %T\n", n)

	// anony functions
	fmt.Println(">>>>> Anony functions")
	fp := pow()
	fmt.Println("pow:", fp())
	fmt.Println("pow:", fp())
	fmt.Println("pow:", fp())
	fmt.Println("pow:", fp())

	// variadic parm
	fmt.Println(">>>>> Variadic params")
	foo(1, 2, 3, 4, 5)

	// send slice variadic
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	foo(xi...)
}

// simple declaration
func multi(n, m int) int {
	return n * m
}

// multiple return value
func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange
}

// return error
func getStockPriceChangeWithError(prevPrice, currentPrice float64) (float64, float64, error) {

	if prevPrice == 0 {
		err := errors.New("error: previous price cannot be zero")
		return 0, 0, err
	}

	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange, nil
}

// named return values (no idiomatic)
func getStockPriceChangeWithNamedReturnValues(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice
	percentChange = (change / prevPrice) * 100
	return
}

// ignore parameter
func sum(_, y, z int) int {
	return y * z
}

// signature
func show(m string) {
	fmt.Println("messague:", m)
}

func product(a, b int) int {
	return a * b
}

func nothing() {
	fmt.Println("do anything")
}

// clouser (clausura o encerramiento)
// return a anony func + a copy scope
func pow() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// variadic parameter
func foo(x ...int) {
	// Go converts x into slice of int
	fmt.Printf("variadic parm, type %T\n", x)
	fmt.Println("variadic", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("Sum of variadic parms:", sum)
}
