package main

import "fmt"

const (
	toKm = 1.61
	toml = 0.62
)

// type
type kilometer float64
type mile float64

func main() {
	// millas y kilometros
	var km float64 = 5.6
	var ml float64 = 3.5
	fmt.Printf("%.2f kms are equal to %.2f miles\n", km, miles(km))
	fmt.Printf("%.2f miles are equal to %.2f kms\n", ml, kilometers(ml))
	// sum
	fmt.Println(km + ml) // logical error
	var k kilometer = 5.6
	var m mile = 3.5
	fmt.Printf("%.2f kms are equal to %.2f miles\n", k, toMiles(k))
	fmt.Printf("%.2f miles are equal to %.2f kms\n", m, tokilometers(m))
	//>> fmt.Println(k + m) // invalid: mismatched types kilometer and mile
	fmt.Println(k + tokilometers(m)) // m was converted in kilometer
}

func miles(k float64) float64 {
	return k * toml
}

func kilometers(m float64) float64 {
	return m * toKm
}

func toMiles(k kilometer) mile {
	return mile(k * toml) // conversion
}

func tokilometers(m mile) kilometer {
	return kilometer(m * toKm) // conversion
}
