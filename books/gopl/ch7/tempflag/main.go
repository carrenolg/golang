package main

import (
	"flag"
	"fmt"
)

// interface
/*type Value interface {
	String()
	Set(string) error
}*/

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

// method
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type celsiusFlag struct {
	Celsius
}

// celsiusFlag: satisfacy the flag.Value interface
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FtoC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// function implement interface
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	c := celsiusFlag{
		value,
	}
	// variable "c" of type celsiusFlag "is a" flag.Value (interface)
	flag.CommandLine.Var(&c, name, usage) //
	return &c.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	// tempconv
	flag.Parse()
	fmt.Println(*temp)

}
