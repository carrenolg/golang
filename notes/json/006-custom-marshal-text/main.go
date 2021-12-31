package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Size int

const (
	Unrecognized Size = iota
	Small
	Large
)

func (s *Size) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	case "small":
		*s = Small
	case "large":
		*s = Large
	default:
		*s = Unrecognized
	}
	return nil
}

func (s Size) MarshalText() ([]byte, error) {
	var name string
	switch s {
	case Small:
		name = "small"
	case Large:
		name = "large"
	default:
		name = "unrecognized"
	}
	return []byte(name), nil
}

func main() {
	blob := `["small","regular","large","unrecognized","small","normal","small","large"]`
	var inventory []Size
	if err := json.Unmarshal([]byte(blob), &inventory); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(inventory)
	counts := make(map[Size]int)
	for _, size := range inventory {
		counts[size] += 1
	}
	fmt.Printf("Inventory Counts:\n*Unrecognized: \t%d\n*Small: \t%d\n*Large: \t%d\n", counts[Unrecognized], counts[Small], counts[Large])
}
