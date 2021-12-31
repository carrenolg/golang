package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	default:
		*a = Unknown
	}
	return nil
}

// Animal implements "MarshalJSON"
func (a Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}
	return json.Marshal(s)
}

func main() {
	// fmt.Println(Unknown, Gopher, Zebra)
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	//data := []byte(blob)
	//fmt.Printf("%T %T\n", blob, data)
	var zoo []Animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}
	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal] += 1
	}
	fmt.Println(census)
	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras: %d\n* UNknowm: %d\n", census[Gopher], census[Zebra], census[Unknown])
}
