package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) // conversion
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9) // conversion
}

func main() {
	var t1 Celsius = 50.0
	var t2 Fahrenheit = 315.6
	fmt.Printf("%2.f°F, %2.f°C\n", CtoF(t1), FtoC(t2))
	//
}
